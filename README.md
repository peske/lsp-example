# What?

This is a demo project with very simple [`github.com/peske/lsp`](https://github.com/peske/lsp) implementation. The
repository contains two projects:

- LSP server written in Go that uses the module mentioned above. It resides in [./server](./server) directory.
- Visual Studio Code extension / LSP client in [./vs-client](./vs-client/) directory. The code is based on
  [`lsp-sample` project](https://github.com/microsoft/vscode-extension-samples/tree/main/lsp-sample) from
  https://github.com/microsoft/vscode-extension-samples repository. We have only _connected_ it to our server.
