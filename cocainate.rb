# inspired by https://github.com/go-task/homebrew-tap/blob/master/Formula/go-task.rb
class Cocainate < Formula
	desc "Cross-platform caffeinate alternative. Pre-compiled."
	homepage "https://github.com/AppleGamer22/cocainate"
	license "GPL-3.0-only"
	head "https://github.com/AppleGamer22/cocainate.git", branch: "master"
	version "1.0.0"

	depends_on "bash" => :optional
	depends_on "fish" => :optional
	depends_on "zsh" => :optional
	depends_on "go" => :build
	depends_on "git" => :build

	on_macos do
		if Hardware::CPU.intel?
			url "https://github.com/AppleGamer22/cocainate/releases/download/#{version}/cocainate_#{version}_mac_amd64.tar.gz"
		end

		if Hardware::CPU.arm?
			url "https://github.com/AppleGamer22/cocainate/releases/download/#{version}/cocainate_#{version}_mac_arm64.tar.gz"
		end
	end

	on_linux do
		depends_on "dbus"

		if Hardware::CPU.intel?
			url "https://github.com/AppleGamer22/cocainate/releases/download/#{version}/cocainate_#{version}_linux_amd64.tar.gz"
		end

		if Hardware::CPU.arm?
			url "https://github.com/AppleGamer22/cocainate/releases/download/#{version}/cocainate_#{version}_linux_arm64.tar.gz"
		end
	end

	def install
		bin.install "cocainate"
		man1.install "cocainate.1"
		bash_completion.install "cocainate.bash" => "cocainate"
		fish_completion.install "cocainate.fish"
		zsh_completion.install "cocainate.zsh" => "_cocainate"
	end

	test do
		system "#{bin}/cocainate", "--help"
	end
end