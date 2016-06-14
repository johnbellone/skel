require 'etc'

knife[:vault_mode] = 'client'
#knife[:vault_admins] = %w{}

chef_server_url 'https://manage.chef.io/organizations/{{ .Org }'
node_name Etc.getlogin
client_key File.join(Dir.home, '.chef', "#{Etc.getlogin}.pem")
