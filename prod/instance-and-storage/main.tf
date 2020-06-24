


module "instance_and_storage" {
	source = "../../../modules/instance-and-storage"
  
	# Compartment
	compartment_ocid = "ocid1.compartment.oc1..aaaaaaaaezxcenaj36sx2s4upu4n76ptsa3mqrkm5ppnu2fswqdpduqywz7q"

	# Compute Instance Configurations
	instance_display_name = "terraform-instance-prod"
	source_ocid = "ocid1.image.oc1.iad.aaaaaaaazvlnvaprv65ak2fqhxzpf6wda2vpbnktiroua3fzrxeizfxzphca"	
	subnet_ocids = ["ocid1.subnet.oc1.iad.aaaaaaaawnby4ouxffhab45r3qjq2ukbffnmi3lkpi2sabwwvxirl3q7lsca"]
	ssh_authorized_keys = "C:\\Users\\matias.araya.cohen\\OneDrive - Accenture\\Documents\\SSH Keys\\public-key"

	# Storage Volume Configurations
	block_storage_sizes_in_gbs = [50]

	compute_availability_domain_list = ["MmBO:US-ASHBURN-AD-1", "MmBO:US-ASHBURN-AD-2", "MmBO:US-ASHBURN-AD-3"]

}



