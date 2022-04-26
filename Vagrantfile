Vagrant.configure("2") do |config|
	config.ssh.insert_key = false
	config.vm.provider "virtualbox" do |virtualbox|
		virtualbox.memory = 2048
		virtualbox.cpus = 2
		virtualbox.gui = false
	end
	
	config.vm.define "ubuntu" do |ubuntu|
		ubuntu.vm.box = "generic/ubuntu2004"
		ubuntu.vm.hostname = "ubuntu"
		ubuntu.vm.synced_folder ".", "/home/vagrant/Documents/cocainate", create: true
		ubuntu.vm.provision "shell", inline: <<-SCRIPT
			# sudo apt update && sudo apt upgrade -y
			# sudo apt install -y ubuntu-desktop-minimal
			# sudo systemctl start gdm3
		SCRIPT
	end
	
	config.vm.define "freebsd" do |freebsd|
		freebsd.vm.box = "generic/freebsd13"
		freebsd.vm.hostname = "freebsd"
		freebsd.vm.synced_folder ".", "/home/vagrant/Documents/cocainate", create: true
	end
end