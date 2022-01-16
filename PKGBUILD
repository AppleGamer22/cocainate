# Maintainer: Omri Bornstein <omribor@gmail.com>
pkgname=cocainate-bin
pkgver=1.0.0
pkgrel=1
pkgdesc="Cross-platform caffeinate alternative. Pre-compiled."
arch=('x86_64' 'aarch64')
url="https://github.com/AppleGamer22/cocainate"
license=('GPL3')
depends=('dbus')

makedepends=(
	'go'
	'git'
)
optdepends=(
	'bash'
	'fish'
	'zsh'
)

provides=('cocainate')
conflicts=('cocainate')

source_x86_64=("https://github.com/AppleGamer22/cocainate/releases/download/${pkgver}/cocainate_${pkgver}_linux_amd64.tar.gz")
source_aarch64=("https://github.com/AppleGamer22/cocainate/releases/download/${pkgver}/cocainate_${pkgver}_linux_arm64.tar.gz")

sha256sums_x86_64=()
sha256sums_aarch64=()

package() {
	# Binary
	install -Dm755 cocainate "${pkgdir}/usr/bin/cocainate"
	# Manual Page
	install -Dm644 cocainate.1 "${pkgdir}/usr/share/man/man1/cocainate.1"

	# Shell Aautocompletion
	install -Dm644 cocainate.bash "${pkgdir}/usr/share/bash-completion/completions/cocainate"
	install -Dm644 cocainate.fish "${pkgdir}/usr/share/fish/vendor_completions.d/cocainate.fish"
	install -Dm644 cocainate.zsh "${pkgdir}/usr/share/zsh/site-functions/_cocainate"
}