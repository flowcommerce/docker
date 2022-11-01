def run(command, quiet=false, fail=true)
  if quiet == false
    puts "  %s" % command
  end
  result = system(command)
  if result != true
    raise "COMMAND FAILED:  %s" % command
  end
  result
end
