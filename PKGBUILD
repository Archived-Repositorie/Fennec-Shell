# Maintainer: JustFoxxo <justafoxxo@outlook.com>
pkgname=fennec-shell
pkgver=1.0
pkgrel=1
epoch=
pkgdesc="A Shell written in Go"
arch=(x86_64)
url="https://github.com/JustFoxx/fennec-shell.git"
license=('GPL')
groups=()
depends=(go)
makedepends=(go)
checkdepends=()
optdepends=(nodejs python lua bash php php-cgi ruby perl powershell)
conflicts=()
replaces=()
backup=()
options=()
install=
changelog=
source=("git+$url")
noextract=()
md5sums=('SKIP')
validpgpkeys=()

pkgver() {
  cd "${_pkgname}"
  printf "1.0.r%s.%s" "$(git rev-list --count HEAD)" "$(git rev-parse --short HEAD)"
}


build() {
	cd "${_pkgname}"
	go build
}

package() {
	cd "${_pkgname}"
	export GOPATH=/
	go install
	fs -install
}
