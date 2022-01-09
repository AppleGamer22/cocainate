# Maintainer: Omri Bornstein <omribor at gmail dot com>
pkgname=cocainate-bin
pkgver=1.0.0
pkgrel=1
pkgdesc="Cross-platform caffeinate alternative. Pre-compiled."
arch=('x86_64' 'aarch64')
url="https://github.com/AppleGamer22/cocainate"
license=('GPL3')
depends=('dbus')
provides=('cocainate')
conflicts=('cocainate')

source_x86_64=("https://github.com/AppleGamer22/cocainate/releases/download/${pkgver}/${pkgname}_${pkgver}_linux_amd64.tar.gz")
source_aarch64=("https://github.com/AppleGamer22/cocainate/releases/download/${pkgver}/${pkgname}_${pkgver}_linux_arm64.tar.gz")

# sha256sums_x86_64=('d15be2b446b19fcf2677c3d14f524f930738054aa30d96beef63ed5481d70321')
# sha256sums_aarch64=('38c720bd3a84cbb29f42b4a48aee31d27b687bcc7658a4769861022790ea5d8e')
# sha256sums_armv6h=('60d7f8fa69f8283d9b2561ddce4667f7a2af1bf06099b1e0fc8bf5da394f626a')
# sha256sums_armv7h=('43b6f2c525f07335f85f90b2925c14d9f720992a2f7291d0d07911d89f16796b')

package() {
	_output="${srcdir}/${pkgname}_${pkgver}_${CARCH}"
	install -Dm755 "${_output}/${pkgname}" "${pkgdir}/usr/bin/${pkgname}"
	# install -Dm644 "${_output}/yay.8" "${pkgdir}/usr/share/man/man8/yay.8"

	# Shell autocompletion script
	# install -Dm644 "${_output}/bash" "${pkgdir}/usr/share/bash-completion/completions/yay"
	# install -Dm644 "${_output}/zsh" "${pkgdir}/usr/share/zsh/site-functions/_yay"
	# install -Dm644 "${_output}/fish" "${pkgdir}/usr/share/fish/vendor_completions.d/yay.fish"
}