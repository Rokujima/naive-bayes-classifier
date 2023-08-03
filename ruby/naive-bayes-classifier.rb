class NaiveBayesClassifier
    def initialize
      @classes = {}
      @total_count = 0
    end
  
    def train(data, target_class)
      @classes[target_class] ||= { count: 0, features: {} }
  
      target_class_data = @classes[target_class]
  
      data.each_with_index do |feature, index|
        target_class_data[:features][index] ||= {}
        target_class_data[:features][index][feature] ||= 0
        target_class_data[:features][index][feature] += 1
      end
  
      target_class_data[:count] += 1
      @total_count += 1
    end
  
    def classify(data)
      max_prob = -Float::INFINITY
      best_class = nil
  
      @classes.each do |target_class, target_class_data|
        class_prob = Math.log(target_class_data[:count].to_f / @total_count)
  
        data.each_with_index do |feature, index|
          feature_data = target_class_data[:features][index] || {}
          feature_count = feature_data[feature] || 0
          total_feature_count = feature_data.values.sum
  
          smoothed_prob = (feature_count + 1).to_f / (total_feature_count + feature_data.size)
  
          class_prob += Math.log(smoothed_prob)
        end
  
        if class_prob > max_prob
          max_prob = class_prob
          best_class = target_class
        end
      end
  
      best_class
    end
  end