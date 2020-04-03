```ruby
# frozen_string_literal: true

require 'aws-sdk'

class FileUploadService # rubocop:disable Metrics/ClassLength
  attr_reader :shard_size, :overwrite, :resource

  # example log option: options[:logger] = DMLogger.aws_logger options[:log_level] = 'info'
  def initialize(options = {})
    options.assert_valid_keys(:region, :endpoint, :force_path_style, :http_wire_trace, :logger,
                              :log_level, :access_key_id, :secret_access_key)
    if options[:access_key_id].present? && options[:secret_access_key].present?
      options[:credentials] = Aws::Credentials.new(options[:access_key_id], options[:secret_access_key])
    end
    options[:region] ||= 'us-west-1'

    set_shard_size(15)
    @overwrite = false
    @check_exists = true
    @region = options[:region]
    @resource = Aws::S3::Resource.new(options)
  end

  # options: {file_path: 'file_path'}
  def upload_file(bucket, key, options = {})
    file_path = options[:file_path]
    file_name = File.basename(options[:file_path])
    options[:file_name] ||= file_name
    options = default_options(options)
    obj = check_key(bucket, key)
    options.except!(:file_name, :file_path, :content)
    obj.upload_file(file_path, options.merge(multipart_threshold: shard_size))
    {region: @region, bucket: bucket, key: key}
  end

  # options: {content: 'content', file_name: 'file_name'}
  def upload_content(bucket, key, options = {})
    content_size = ObjectSpace.memsize_of(options[:content])
    raise "Content size is bigger than #{shard_size}" if content_size > shard_size

    content = options[:content]
    options = default_options(options)
    obj = check_key(bucket, key)
    options.except!(:file_name, :file_path, :content)
    obj.put(body: content, **options)
    {region: @region, bucket: obj.bucket.name, key: key}
  end

  # options: {content: 'content', file_name: 'file_name'}
  def multipart_upload_content(bucket, key, options = {})
    content_size = ObjectSpace.memsize_of(options[:content])
    raise "Content size is smaller than #{shard_size}" if content_size < shard_size

    options = default_options(options)
    check_key(bucket, key)
    dd = StringIO.new(options[:content])

    files = split_file(dd)
    multi_upload_to_s3(bucket, key, files, options)
  end

  # max expire_in is one week
  def get_presigned_url(bucket, key, expire_in = 3600, version_id = nil)
    Aws::S3::Object.new(key: key, bucket_name: bucket, client: resource.client).presigned_url(
      :get, expires_in: expire_in,
            version_id: version_id
    )
  end

  def shard_size=(chunk_size)
    @shard_size = chunk_size * 1024 * 1024
  end
  alias_method :set_shard_size, :shard_size=

  def overwrite=(is_overwrite)
    @overwrite = is_overwrite
  end

  def check_exists=(if_check)
    @check_exists = if_check
  end

  private

  def check_key(bucket, key)
    bucket = @resource.bucket(bucket)
    obj = bucket.object(key)
    raise 'The key is already exist' if @check_exists &&
                                        obj.exists? &&
                                        !@overwrite

    obj
  end

  def default_options(options)
    options.assert_valid_keys(:content_disposition, :content_type, :acl, :file_name, :content, :file_path)
    options[:content_disposition] ||= options[:file_name] ? "attachment; filename=#{options[:file_name]}" : 'attachment'
    options[:content_type] ||= 'application/octet-stream'
    options[:acl] ||= 'private'
    options
  end

  def split_file(file)
    temp_name = SecureRandom.uuid
    files = []
    until file.eof?
      part_number = format('%05d', (file.pos / shard_size + 1))
      temp = File.new(temp_name + "-part-#{part_number}", 'w')
      temp.binmode
      content = file.read(shard_size)
      temp.write(content)

      files << {
        path: temp.path,
        size: content.bytesize
      }
    end
    file.close unless files.empty?
    files
  end

  def multi_upload_to_s3(bucket, key, files, options = {}) # rubocop:disable Metrics/AbcSize
    options.except!(:file_name, :file_path, :content)
    client = resource.client
    parts = client.create_multipart_upload(
      bucket: bucket,
      key: key,
      **options
    )
    # {:bucket=>"", :key=>"", :upload_id=>""}
    opts = parts.to_h
    result = opts.merge(multipart_upload: { parts: [] })

    result[:multipart_upload][:parts] =
      Parallel.map(1..files.size) do |part_number|
        file_info = files[part_number - 1]
        file = File.open(file_info[:path], 'rb')

        response = client.upload_part(
          opts.merge(body: file.read, part_number: part_number, content_length: file_info[:size])
        )

        file.close
        response.to_h.merge(part_number: part_number)
      end

    client.complete_multipart_upload(result)

    files.map {|file| File.delete(file[:path]) if File.exist?(file[:path])}
    {region: @region, bucket: bucket, key: key}
  end
end
```