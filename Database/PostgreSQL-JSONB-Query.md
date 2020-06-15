```ruby
# payload: [{"kind"=>"person"}]
Segment.where("payload @> ?", [{kind: "person"}].to_json)

# data: {"interest"=>["music", "movies", "programming"]}
Segment.where("data @> ?",  {"interest": ["music", "movies", "programming"]}.to_json)
Segment.where("data #>> '{interest, 1}' = 'movies' ")
Segment.where("jsonb_array_length(data->'interest') > 1")
Segment.where("data->'interest' ? :value", value: "movies") 
Segment.where("data -> 'interest' ? :value", value: ['programming'])

# data: {"customers"=>[{:name=>"david"}]}
Segment.where("data #> '{customers,0}' ->> 'name' = 'david' ")
Segment.where("data @> ?",  {"customers": [{"name": "david"}]}.to_json)
Segment.where("data -> 'customers' @> '[{\"name\": \"david\"}]'")
Segment.where(" data -> 'customers' @> ?", [{name: "david"}].to_json)

# data: {"uid"=>"5", "blog"=>"recode"}
Segment.where("data @> ?", {uid: '5'}.to_json)
Segment.where("data ->> 'blog' = 'recode'")
Segment.where("data ->> 'blog' = ?", "recode")
Segment.where("data ? :key", :key => 'uid')
Segment.where("data -> :key LIKE :value", :key => 'blog', :value => "%recode%")

# tags: ["dele, jones", "solomon"]
# get a single tag
Segment.where("'solomon' = ANY (tags)")
# which segments are tagged with 'solomon'
Segment.where('? = ANY (tags)', 'solomon')
# which segments are not tagged with 'solomon'
Segment.where('? != ALL (tags)', 'solomon')
Segment.where('NOT (? = ANY (tags))', 'solomon')
# multiple tags
Segment.where("tags @> ARRAY[?]::varchar[]", ["dele, jones", "solomon"])
# tags with 3 items
Segment.where("array_length(tags, 1) >= 3")
```