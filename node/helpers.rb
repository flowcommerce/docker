def run(command, quiet=false, fail=true)
  if quiet == false
    puts "  %s" % command
  end
  if fail == false
    puts "exception rescued"
  result = system(command)
  if result != true
    raise "COMMAND FAILED:  %s" % command
  end
  result
end
