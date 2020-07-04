class AnimalShelf
    def initialize()
      @cats = []
      @dogs = []
    end


=begin
    :type animal: Integer[]
    :rtype: Void
=end
    def enqueue(animal)
      if animal[1] == 0
        @cats.push(animal)
      else
        @dogs.push(animal)
      end
    end


=begin
    :rtype: Integer[]
=end
    def dequeue_any()
      case true
      when @cats.empty? || @dogs.empty?
        [-1, -1]
      when @cats.empty? # || !@dogs.empty?
        @dogs.shift
      when @dogs.empty?
        @cats.shift
      else
        @cats[0][0] < @dogs[0][0] ? @cats.shift : @dogs.shift
      end
    end


=begin
    :rtype: Integer[]
=end
    def dequeue_dog()
      @dogs.shift || [-1,-1]
    end


=begin
    :rtype: Integer[]
=end
    def dequeue_cat()
      @cats.shift|| [-1,-1]
    end


end

# Your AnimalShelf object will be instantiated and called as such:
# obj = AnimalShelf.new()
# obj.enqueue(animal)
# param_2 = obj.dequeue_any()
# param_3 = obj.dequeue_dog()
# param_4 = obj.dequeue_cat()