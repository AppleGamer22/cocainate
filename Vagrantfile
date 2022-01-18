Vagrant.configure("2") do |config|
	config.ssh.insert_key = false
	config.vm.provider "virtualbox" do |virtualbox|
		virtualbox.memory = 2048
		virtualbox.cpus = 2
		virtualbox.gui = false
		virtualbox.name = "cocainate"
	end
	config.vm.box = "ubuntu/focal64"
	config.vm.hostname = "ubuntu"
	config.vm.synced_folder ".", "/home/vagrant/Documents/cocainate", create: true
	config.vm.provision "shell", inline: <<-SCRIPT
		sudo apt update && sudo apt upgrade -y
		sudo apt install -y ubuntu-desktop golang
		sudo systemctl start gdm3
	SCRIPT
end