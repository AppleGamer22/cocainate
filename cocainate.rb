class Cocainate < Formula
	desc "Cross-platform caffeinate alternative. Pre-compiled."
	homepage "https://github.com/AppleGamer22/cocainate"
	url "https://example.com/foo-0.1.tar.gz"
	sha256 "85cc828a96735bdafcf29eb6291ca91bac846579bcef7308536e0c875d6c81d7"
	license "GPL-3.0-only"
	head "https://github.com/AppleGamer22/cocainate.git", branch: "master"
	version "1.0.0"

	depends_on "bash" => :optional
	depends_on "fish" => :optional
	depends_on "zsh" => :optional

	on_linux do
		depends_on "dbus"
	end

	def install
		# system "ls"
		# #{prefix}/bin
		# #{prefix}/share/man
	end

	test do
		assert_equal version, shell_output("#{prefix}/bin/cocainate version").strip
	end
end