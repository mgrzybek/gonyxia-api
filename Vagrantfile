# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = ENV["VAGRANT_BOX_NAME"] || "ubuntu/jammy64"
  config.vm.box_url = ENV["VAGRANT_BOX_URL"]
  config.vm.hostname = "gonyxia-api-k3s"

  config.vm.provider "virtualbox" do |p|
    p.cpus = ENV['VAGRANT_VM_CPUS'] || 1
    p.memory = ENV['VAGRANT_VM_MEMORY'] || 1024
  end

  config.vm.provision "ansible" do |ansible|
    ansible.playbook = "tests/main.yml"
  end
end
