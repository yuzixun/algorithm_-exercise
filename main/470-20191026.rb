# The rand7() API is already defined for you.
# def rand7()
# @return {Integer} a random integer in the range 1 to 7

def rand10()
  result = 40
  while result >= 40
    result = (rand7() - 1) * 7 + (rand7() - 1)
  end
  result%10+1
end