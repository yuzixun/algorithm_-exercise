# @param {String} text
# @return {String}
def entity_parser(text)
  text.gsub("&quot;", "\"").gsub("&apos;", "\'").gsub("&gt;", ">" ).gsub("&lt;", "<" ).gsub("&frasl;", "/" ).gsub("&amp;", "&")
end

p entity_parser("&amp; is an HTML entity but &ambassador; is not.")
p entity_parser("and I quote: &quot;...&quot;")
p entity_parser("Stay home! Practice on Leetcode :)")
p entity_parser("x &gt; y &amp;&amp; x &lt; y is always false")
p entity_parser("leetcode.com&frasl;problemset&frasl;all")