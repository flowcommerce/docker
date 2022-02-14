def run(command, quiet=false)
  if quiet == false
    puts "  %s" % command
  end
  result = system(command)
  if result != true
    raise "COMMAND FAILED:  %s" % command
  end
  result
end

def get_env_providor_location()
  possible_paths = [
    '/../../environment-provider',
    '/../../../flowvault/environment-provider'
  ];
  expanded_paths = possible_paths.map { |path| File.expand_path(File.join(File.dirname(__FILE__), path)) }
  found_path = expanded_paths.find { |path| File.directory?(path) }

  if found_path == nil
    puts "ERROR finding environment-provider repo"
    exit(1)
  end

  puts "Found environment-provider repo at: [#{found_path}]"
  found_path
end
