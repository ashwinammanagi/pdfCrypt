# 📄 pdfCrypt

**pdfCrypt** is a fast, secure, and cross-platform desktop application to **encrypt** and **decrypt PDF files** using a clean and user-friendly UI built with [Wails](https://wails.io) and Svelte.

### 🛡️ Why this app?

Most free tools for PDF encryption or decryption are either:

- **Windows-only**, or
- **Online services** where you must upload sensitive documents — something we don’t feel safe doing.

**pdfCrypt** runs **entirely on your device**, with no internet dependency. It's built with [pdfcpu](https://github.com/pdfcpu/pdfcpu) under the hood and works on **macOS**, **Windows**, and **Linux**.

---

## ✨ Features

- 🔒 **Encrypt PDFs** with a password
- 🔓 **Decrypt PDFs** using the correct password
- 🎛️ Set PDF permissions: print, modify, extract, annotate, and more
- ⚡ Clean UI powered by Svelte
- 💯 Offline-first and open source

---

## 📥 Download

Pre-built binaries for **macOS**, **Windows**, and **Linux** will be available on the [Releases](https://github.com/ashwinammanagi/pdfCrypt/releases) page.

---

## 🚀 Getting Started (for Developers)

### Prerequisites

- [Go](https://go.dev/doc/install)
- [Node.js + pnpm](https://pnpm.io/)
- [Wails CLI](https://wails.io/docs/gettingstarted/installation)

Install Wails if not already installed:

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

---

## 💻 Run in Development

To start the app in development mode with hot reload:

```bash
wails dev
```

This will:

- Start the backend Go server
- Launch Vite for frontend hot-reloading

---

## 📦 Build for Production

To build a standalone desktop app:

```bash
wails build
```

Binaries will be generated in the `build/bin/` directory.

### 🔧 Building for Specific Platforms

You can cross-compile from any OS using the `-platform` flag:

```bash
# Build for macOS
wails build -platform darwin/amd64

# Build for Windows
wails build -platform windows/amd64

# Build for Linux
wails build -platform linux/amd64
```

For cross-compilation to work smoothly, install the required OS SDKs or Docker images as [per Wails documentation](https://wails.io/docs/guides/crossplatform).

---

## 🛠 Usage

1. Launch `pdfCrypt`
2. Choose **Encrypt** or **Decrypt**
3. Click **Choose PDF** to select a file
4. Enter a **password**
5. (If encrypting) Select desired permissions
6. Click the button to encrypt or decrypt

You’ll see a message with the output file path once it’s done.

---

## 🔐 Permissions Explained

When encrypting, you can control what users are allowed to do with the PDF:

- ✅ Print
- ✏️ Modify
- 📋 Extract content
- 🖊️ Annotate
- 📝 Fill forms
- 📚 Assemble
- 🖨️ Print high resolution

These are enforced based on the PDF standard, though readers may vary in how strictly they apply them.

---

## 📄 License

This project is licensed under the MIT License. See [LICENSE](./LICENSE) for more.

---

## 🙏 Acknowledgements

- [pdfcpu](https://github.com/pdfcpu/pdfcpu) for powerful PDF processing
- [Wails](https://wails.io) for modern desktop app development with Go
- [Svelte](https://svelte.dev) for a super clean frontend experience

---

## 📬 Contributions

Pull requests and ideas welcome! Please open an issue to discuss features or bugs.
