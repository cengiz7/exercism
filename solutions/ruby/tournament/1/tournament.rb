module Tournament
    def self.tally(input)
        @@scores = Hash.new { |h, k| h[k] = {MP: 0, W: 0, D: 0, L: 0, P: 0} }

        input.strip.each_line do |line|
            team1, team2, result = line.strip.split(';')
            add_team_score(team1, result)
            add_team_score(team2, get_opponent_result(result))
        end
        draw_table
    end

    def self.add_team_score(team,result)
        result, points = get_result_with_point(result)
        @@scores[team][result] += 1
        @@scores[team][:MP] += 1
        @@scores[team][:P] += points
    end

    def self.get_result_with_point(result)
        case result
        when 'win' then [:W, 3]
        when 'draw' then [:D, 1]
        when 'loss' then [:L, 0]
        end
    end

    def self.get_opponent_result(result)
        case result
        when 'win' then 'loss'
        when 'draw' then 'draw'
        when 'loss' then 'win'
        end
    end

    def self.sorted_scores
        @@scores.sort_by {|k, v| [-v[:P], k]}
    end

    def self.draw_table
        table_header = "Team                           | MP |  W |  D |  L |  P\n"
        table_scores = sorted_scores.collect{|team, sc| sprintf("%-30s |%3d |%3d |%3d |%3d |%3d\n",
                                            team, sc[:MP], sc[:W], sc[:D], sc[:L], sc[:P] )
                                        }.join
        table_header + table_scores
    end

    class << self
        private :get_opponent_result, :get_result_with_point,
                :add_team_score, :draw_table, :sorted_scores
    end
end