class Rule

    def initialize(colors)
        @color_list = colors
    end

    def accept_next(color)
    end
end

class AllowCut < Rule

    def accept_next(color)
        @color_list.include?(color)
    end
end

class ForbidCut < Rule

    def accept_next(color)
        !@color_list.include?(color)
    end
end

class RuleSet

    def initialize
        @rules = {w: ForbidCut.new( [:w, :b]),
                  r: AllowCut.new([:g]),
                  b: ForbidCut.new([:w,:g,:o]),
                  o: AllowCut.new([:r,:b]),
                  g: AllowCut.new([:o,:w]),
                  p: ForbidCut.new([:p,:g,:o,:w])}
    end

    def accept_sequence(color_sequence)
        (color_sequence.length- 1).times { |i| 
            if !@rules[color_sequence[i].to_sym].accept_next(
                    color_sequence[i+1].to_sym)
                puts "Boom"
                return
            end
        }
        puts "Bomb defused"
    end
end

rules = RuleSet.new
rules.accept_sequence( ["w","r","g","w"])
rules.accept_sequence( ["w","o","g","w"])
