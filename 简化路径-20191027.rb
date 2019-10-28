# @param {String} path
# @return {String}
def simplify_path(path)
  size = path.size

  pathes = path.split("/").delete_if { |c| c == "." || c == "" }

  result = []

  pathes.each do |p|
    if p == ".."
      result.pop
    else
      result.push(p)
    end
  end

  "/" + result.join("/")
end


# s = "/a//b////c/d//././/.."
# s = "/../"
s = "/home//foo/"
puts simplify_path(s)