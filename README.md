# TTMS

# 修改远程仓库名，区别于github远程仓库
```bash
git remote rename
 或
git remote add huawei git@codehub.devcloud.cn-north-4.huaweicloud.com:your-project/TTMS.git
```

推送华为云：
```bash
git push huawei main
密码：lion123
```

推送github
```bash
git push origin main
```

# 在github和华为云上分别使用git上传

需要在 git-bash 的 .ssh 目录下创建 config 文件，并添加如下内容：

```bash
1. 创建 .ssh 目录（如果不存在）
mkdir -p ~/.ssh

2. 创建 config 文件
touch ~/.ssh/config

3. 打开并编辑 config 文件
```

```text
# GitHub 配置
Host github.com
    HostName github.com
    IdentityFile ~/.ssh/id_rsa
    User git

# 华为云配置
Host codehub.devcloud.cn-north-4.huaweicloud.com
    HostName codehub.devcloud.cn-north-4.huaweicloud.com
    IdentityFile C:/Users/11067/.ssh/id_rsa_huaweicloud
    User git
    IdentitiesOnly yes
```

同时，需要给 config 文件添加权限：

```bash
chmod 600 ~/.ssh/config
```

```bash

测试连接：
ssh -T git@codehub.devcloud.cn-north-4.huaweicloud.com

```

# 部署
```bash
// 后台运行
nohup ./main

// 退出后台运行
netstat -anp | grep 8080
kill 进程号
```