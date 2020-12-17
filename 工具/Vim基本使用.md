## Vundle插件管理

---
### 插件
1. preservim/nerdtree
2. vim-ruby/vim-ruby
3. tpope/vim-rails
4. Yggdroot/LeaderF
5. Shougo/neocomplete.vim

---
### 配置
```vim
" Plugin 配置
" nerdtree
autocmd StdinReadPre * let s:std_in=1
autocmd VimEnter * if argc() == 1 && isdirectory(argv()[0]) && !exists("s:std_in") | exe 'NERDTree' argv()[0] | wincmd p | ene | exe 'cd '.argv()[0] | endif

"vim基本配置
"缩进
set tabstop=2
set backspace=2
set autoindent
set softtabstop=4
set shiftwidth=4
" 智能自动缩进
" 设置自动缩进
set smartindent
set ai!
" 显示括号配对情况
set showmatch
set autoindent
set cindent
set sw=2
"show curosr position all the time"
set  ruler
"display incomplete commands
set showcmd
syntax enable
"背景颜色
set background=dark
"行号
set number
"搜索高亮
set hlsearch
syntax on

" Javascript syntax hightlight
syntax enable

" Highlight current line
au WinLeave * set nocursorline nocursorcolumn
au WinEnter * set cursorline cursorcolumn
set cursorline cursorcolumn

" Set syntax highlighting for specific file types
autocmd BufRead,BufNewFile Appraisals set filetype=ruby
autocmd BufRead,BufNewFile *.md set filetype=markdown
autocmd Syntax javascript set syntax=jquery

highlight NonText guibg=#060606
highlight Folded  guibg=#0A0A0A guifg=#9090D0
```

---
### 快捷键
dd：删除一行
G： 行底
gg: 第一行
shfit 4 行尾
shift 6 行首

剪切:
1. 定位鼠标到剪切的开始位置
2. 输入v键开始选择剪切的字符，或者V键是为了选择 整行
3. 移动方向键到结束的地方
4. d键是剪切，y键是复制
5. 移动鼠标到粘贴的位置
6. 输入P是在鼠标位置前粘贴，输入p键是在鼠标的位置后粘贴

**NERDTREE**
在NERDTree操作区的一些基本操作:
```shell
?: 快速帮助文档
o: 打开一个目录或者打开文件，创建的是buffer，也可以用来打开书签
go: 打开一个文件，但是光标仍然留在NERDTree，创建的是buffer
i: 水平分割创建文件的窗口，创建的是buffer
gi: 水平分割创建文件的窗口，但是光标仍然留在NERDTree
s: 垂直分割创建文件的窗口，创建的是buffer
gs: 和gi，go类似
x: 收起当前打开的目录
X: 收起所有打开的目录
e: 以文件管理的方式打开选中的目录
D: 删除书签
P: 大写，跳转到当前根路径
p: 小写，跳转到光标所在的上一级路径
K: 跳转到第一个子路径
J: 跳转到最后一个子路径
C: 将根路径设置为光标所在的目录
u: 设置上级目录为根路径
U: 设置上级目录为跟路径，但是维持原来目录打开的状态
r: 刷新光标所在的目录
R: 刷新当前根路径
I: 显示或者不显示隐藏文件
f: 打开和关闭文件过滤器
q: 关闭NERDTree
A: 全屏显示NERDTree，或者关闭全屏
```

**LEADERF**
:LeaderfFile 查找文件
TAB: 进入选择文件