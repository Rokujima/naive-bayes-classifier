class NaiveBayesClassifier:
    def __init__(self):
        self.classes = {}
        self.total_count = 0

    def train(self, data, target_class):
        if target_class not in self.classes:
            self.classes[target_class] = {
                'count': 0,
                'features': {},
            }

        target_class_data = self.classes[target_class]

        for feature_index, feature in enumerate(data):
            if feature_index not in target_class_data['features']:
                target_class_data['features'][feature_index] = {}

            if feature not in target_class_data['features'][feature_index]:
                target_class_data['features'][feature_index][feature] = 0

            target_class_data['features'][feature_index][feature] += 1

        target_class_data['count'] += 1
        self.total_count += 1

    def classify(self, data):
        max_prob = float("-inf")
        best_class = None

        for target_class, target_class_data in self.classes.items():
            class_prob = target_class_data['count'] / self.total_count

            for feature_index, feature in enumerate(data):
                feature_data = target_class_data['features'].get(feature_index, {})
                feature_count = feature_data.get(feature, 0)
                total_feature_count = sum(feature_data.values())

                smoothed_prob = (feature_count + 1) / (total_feature_count + len(feature_data))

                class_prob *= smoothed_prob

            if class_prob > max_prob:
                max_prob = class_prob
                best_class = target_class

        return best_class

