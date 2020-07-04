# @param {Integer[][]} people
# @return {Integer[][]}
def reconstruct_queue(people)
  result = []
  people_hash = people.group_by(&:first).sort.reverse.to_h

  people_hash.each do |height, people|
    people.sort.each do |person|
      result.insert(person.last, person)
    end
  end

  result
end


# p reconstruct_queue [[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
p (reconstruct_queue [[9,0],[7,0],[1,9],[3,0],[2,7],[5,3],[6,0],[3,4],[6,2],[5,2]]) == [[3,0],[6,0],[7,0],[5,2],[3,4],[5,3],[6,2],[2,7],[9,0],[1,9]]

#