Vagrant.configure("2") do |config|
	config.ssh.insert_key = false
	config.vm.provider "virtualbox" do |virtualbox|
		virtualbox.memory = 2048
		virtualbox.cpus = 2
		virtualbox.gui = false
	end
	
	config.vm.define "ubuntu" do |ubuntu|
		ubuntu.vm.box = "ubuntu/jammy64"
		ubuntu.vm.hostname = "ubuntu"
		ubuntu.vm.synced_folder ".", "/home/vagrant/Documents/cocainate", create: true
		ubuntu.vm.provision "shell", inline: <<-SCRIPT
			sudo apt update
			# sudo apt upgrade -y
			sudo apt install -y kde-plasma-desktop xinit sddm golang
			# sudo systemctl disable xdm
			# sudo systemctl enable --now sddm
			cp /etc/X11/xinit/xinitrc ~/.xinitrc
			startx
		SCRIPT
	end
	
	# config.vm.define "freebsd" do |freebsd|
	# 	freebsd.vm.box = "generic/freebsd13"
	# 	freebsd.vm.hostname = "freebsd"
	# 	freebsd.vm.synced_folder ".", "/home/vagrant/Documents/cocainate", create: true
	# 	freebsd.vm.provision "shell", inline: <<-SCRIPT
	# 		sudo pkg update
	# 		sudo pkg install kde5 sddm xorg
	# 		sysrc dbus_enable="YES" && service dbus start
	# 		sysrc sddm_enable="YES" && service sddm start
	# 	SCRIPT
	# end
end