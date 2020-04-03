## Git安装：
```shell
sudo apt-get install git
yum install git
git或者git --version可以查看相关信息
git config --global core.editor "vim"
git config log.date iso-local
```

## 工作区和缓存区
就是你在电脑里能看到的目录，比如这里的GitHub文件夹就是一个工作区。
![IMAGE](resources/24497EAA269E6429B31ACC3F0A307391.jpg =546x258)
工作区有一个隐藏目录.git，这个不算工作区，而是 Git 的版本库。
Git 的版本库里存了很多东西，其中最重要的就是称为 stage（或者叫 index）的暂存区，还有 Git 为我们自动创建的第一个分支 master，以及指向 master 的一个指针叫 HEAD。
前面讲了我们把文件往 Git 版本库里添加的时候，是分两步执行的：
第一步是用git add把文件添加进去，实际上就是把文件修改添加到暂存区；
第二步是用git commit提交更改，实际上就是把暂存区的所有内容提交到当前分支。
因为我们创建 Git 版本库时，Git 自动为我们创建了唯一一个 master 分支，所以，现在，git commit就是往 master 分支上提交更改。

## Git基本命令
**初始化，添加，提交，推送**
```shell
git init :将这个目录变成git可以管理的仓库。
git add git.txt :将文件添加到暂存区
git add . :将所有文件添加到暂存区
git commit -m "wrote a readme file" :将文件提交到版本库，-m后面是对本次提交的说明
git commit --amend :修改commit信息
git status ：查看当前仓库的状态
git clone [url]:克隆项目,url为你想要复制的项目地址
git clone -b branch remoteaddress :选择分支下载项目
git push origin branch :上传某分支的代码
git pull origin branch :同步某分支代码
```

**查看日志**
```shell
git log :可以查看从最近到最远的提交日志。
git log --pretty=oneline :(查看简略一点)
git log --graph :命令可以看到分支合并图
```

**版本回退**
```shell
在 Git 中，用 HEAD 表示当前版本，上一个版本就是HEAD^，上上一个版本就是HEAD^^，当然往上 100 个版本写 100 个^比较容易数不过来所以写成 HEAD~100。
git reset --hard HEAD^ ：回退到上一个版本
git reset --hard 3628164 ：hard后面是对应版本的commit id
版本号没必要写全，前几位就可以了，Git 会自动去找。当然也不能只写前一两位，因为 Git 可能会找到多个版本号，就无法确定是哪一个了。
现在，你回退到了某个版本，关掉了电脑，第二天早上就后悔了，想恢复到新版本怎么办？找不到g新版本的 commit id 怎么办？在 Git 中，总是有后悔药可以吃的。Git 提供了一个命令
git reflog :用来记录你的每一次命令。

git log -p filename :查看某个文件的详细提交记录.
git blame filename :查看当前每一行是哪个最后提交改动的.
```

**分支管理**
```shell
git branch :查看分支
git branch -a ：查看本地和远程的分支
git branch -r :查看远程分支
git branch mqx_test :创建名为mqx_test的分支
git checkout mqx_test ：切换到mqx_test分之下
git chekcout -b mqx_test :创建并切换分支
git branch -d xxx :删除分支
git branch -m old_branch new_branch ：重命名分支
git brach -m new_branch
git push origin :old-name new-name ：删除远程旧分支推送新分支
git push origin branch:branch ：推送本地分支到远程
git merge branch :合并分支到当前分支,合并分支时，加上 --no-ff 参数就可以用普通模式合并，合并后的历史有分支，能看出来曾经做过合并，而 fast forward 合并就看不出来曾经做过合并。
git merge --squash another: 合并多个提交为一条
git branch -m oldName  newName git push origin newName :重命名分支
git push --delete origin oldName :删除远程分支
```

**备份当前工作区内容**
```shell
git stash :可以把当前工作现场“储藏”起来，等以后恢复现场后继续工作
git stash list :查看保存的工作现场。
git stash apply :恢复，但是恢复后，stash 内容并不删除。
git stash pop :恢复的同时把 stash 内容也删了。默认是最新一个stash
git stash pop stash@{x} :恢复制定的stash并删除 
git stash drop :删除备份
git stash show -p stash@{x} :查看指定stash的修改内容
```

**撤销修改:这里注意命令 git reset --hard, --mixed, --soft 之间的区别**
```shell
git checkout -- readme.txt :把 readme.txt 文件在工作区的修改全部撤销
git reset HEAD file :可以把暂存区的修改撤销掉（unstage），重新放回工作区。
git reset --soft commit_id :回退到某一次提交，保留未提交的修改
1:当你改乱了工作区某个文件的内容，想直接丢弃工作区的修改时
用命令git checkout -- file。
2:当你不但改了工作区某文件内容，还添加到了暂存区，想丢弃修改,分两步，第一步用命令git reset HEAD file，就回到了 1，第二步按 1 操作。
3:将修改的内容commit了，撤销commit，保存修改：git reset --soft commit_id
```

**对比修改**
```shell
git diff :工作区与暂存区
git diff --cached ：暂存区与commit
git diff HEAD :工作区与commit
git diff commit_id file_name :查看某个文件的更改
git diff branch1 branch2 filename :两个分支的某个文件区别
```

**远程相关**
```shell
git remote -v :查看远程仓库地址
git remote add origin git@github.com:miaoqingxin/rails.git :添加远程仓库地址
git remote rm origin :删除远程仓库地址
git remote rename origin ：重命名远程仓库地址
git push -u origin master :这里的 -u 参数选项指定一个默认地址
git fetch :将远程主机的更新全部取回本地
git fetch origin branch_name :取回指定分支的更新.
```

**GIT-REBASE**
1. 牢记: 只对尚未推送或分享给别人的本地修改执行变基操作清理历史，从不对已推送至别处的提交执行变基操作.
2. git rebase 
命令可以用来代替merge来合并代码，修改commit历史。有人将其翻译成变基，也有人翻译成衍合,
on develop: git rebase master: 将develop分支上的提交以补丁的形式，将develop分支上的提交移动到master分支上
3. git rebase --onto master server client: 取出 client 分支，找出处于 client 分支和 server 分支的共同祖先之后的修改，然后把它们在 master 分支上重放一遍
4. 交互式rebase: git rebase -i

```shell
git fetch origin
git rebase origin/master
```

**Cherry-pick**
```shell
# 合并单个commit
git cherry-pick hash值
# 合并多个提交
git cherry-pick 第一个hash值 第二个hash值 第三个hash值
# 合并一个区间
git cherry-pick 开始的hash值..结束的hash值
```

**创建标签**:
相当于版本库的一个快照。
切换到需要打标签的分支上
```shell
git tag <name> :就可以打一个新标签,默认标签是打在最新提交的 commit 上的。  
git tag v0.9 6224937 :打tag到对应的 commit id 是 6224937
git tag :查看标签
git show <tagname> :查看标签信息

git tag -a v0.1 -m "version 0.1 released" 3628164 :创建带有说明的标签，用 -a 指定标签名， -m 指定说明文字
git tag -d v0.1 :删除标签
git push origin <tagname> :推送某个标签到远程  
如果标签已经推送到远程，要删除远程标签就麻烦一点，先从本地删除：
 git tag -d v0.9
然后，从远程删除。删除命令也是 push，但是格式如下：
 git push origin :refs/tags/v0.9
要看看是否真的从远程库删除了标签，可以登陆 GitHub 查看
```

## 忽略文件
不想将某个文件加入版本控制中，可以在项目根目录下新建一个名为 .gitignore 的文件
文件内容有个参考[gitignore](https://github.com/github/gitignore)

## github设置项目语言
rails项目自然最好被github识别weiruby项目，但是若html和css过多则会被识别html
```shell
# gitattributes设置项目语言
*.html linguist-language=JavaScript
```
PR reviewer comment: LGTM: look good to me

## git warning:  CRLF will be replaced by LF in xxx
原因分析：
CRLF -- Carriage-Return Line-Feed 回车换行
就是回车(CR, ASCII 13, \r) 换行(LF, ASCII 10, \n)。
这两个ACSII字符不会在屏幕有任何输出，但在Windows中广泛使用来标识一行的结束。而在Linux/UNIX系统中只有换行符。
也就是说在windows中的换行符为 CRLF， 而在linux下的换行符为：LF

## mac:error: There was a problem with the editor 'vi'
```shell
git config --global core.editor /usr/bin/vim
```
## Git ignore文件中添加了某文件，但是 git status 还是出现了改文件
```shell
git rm --cached file
git commit
```

## Git Submodule
```shell
git submodule add submodule_url
git submodule init
git submodule update
```

## 工具推介
[tig](https://github.com/jonas/tig)
[git-extras](https://github.com/tj/git-extras)