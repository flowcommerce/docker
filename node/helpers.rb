def run(command, quiet=false, fail=true)
  if quiet == false
    puts "  %s" % command
  end
  begin
    if fail == false
  rescue
    puts "exception rescued"
  end
  result = system(command)
  if result != true
    raise "COMMAND FAILED:  %s" % command
  end
  result
end
