### Dockerfile
---
docker build -t miao/sugon:v1 .
```docker
FROM rtyler/rvm:latest

RUN sudo apt update && sudo apt install -y libxslt-dev libxml2-dev libmysqlclient-dev wget zsh tig imagemagick

RUN /bin/bash -l -c "rvm install 1.9.3"

RUN /bin/bash -l -c "rvm use 1.9.3 --default"

RUN /bin/bash -l -c "gem install bundler -v 1.7.2"

RUN git clone https://github.com/ohmyzsh/ohmyzsh.git ~/.oh-my-zsh
RUN cp ~/.oh-my-zsh/templates/zshrc.zsh-template ~/.zshrc
COPY .zshrc ~/.zshrc
```

1. RVM需要以登录shell的方式访问，无法直接使用
2. 因为网络问题，oh-my-zsh的安装采用clone的方式

### 中文问题
---
```
locale
locale -a
export LC_ALL="C.UTF-8"
```

### docker export and save
---
docker export 导出container，import之后会丢失CMD等信息，可以通过添加 --change参数来添加到镜像中
dokeer save 保存image，docker load 加载镜像
```shell
cat sugon.tar | docker import --change "CMD ["/bin/bash"]"  - test/sugon:1.0
docker export container-id > export.tar
docker save -o sugon.tar image-name
docker
```

### COMMANDS
---
```shell
docker container rename container new-name
docker system/image/volumn/container prune
```

### 复制文件
---
```shell
docker cp <containerId>:/file/path/within/container /host/path/target
```