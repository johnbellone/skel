default_source :chef_repo, File.expand_path('../..', .)
default_source :community

run_list 'poise-hoist::default'
