def run(command, quiet=false, fail=true)
 begin
  if quiet == false
    puts "  %s" % command
  end
  if fail == false
    puts ""
  end
 rescue
   puts "exception rescued"
 end
  result = system(command)
  if result != true
    raise "COMMAND FAILED:  %s" % command
  end
  result
end
