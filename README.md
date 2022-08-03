# Vue + Go

简单的桌面工具应用程序

## 开发环境

使用前需安装 Node.js，Go

随后安装 pnpm，wails

```shell
iwr https://get.pnpm.io/install.ps1 -useb | iex
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

## 构建

```shell
pnpm install
pnpm build
```

## 启动

```shell
start ./toolkit.exe
```

## 开发预览

```shell
pnpm dev
```
