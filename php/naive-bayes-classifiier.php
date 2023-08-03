<?php
class NaiveBayesClassifier {
    private $classes = [];
    private $totalCount = 0;

    public function train($data, $targetClass) {
        if (!isset($this->classes[$targetClass])) {
            $this->classes[$targetClass] = [
                'count' => 0,
                'features' => [],
            ];
        }

        $targetClassData = &$this->classes[$targetClass];

        foreach ($data as $index => $feature) {
            if (!isset($targetClassData['features'][$index])) {
                $targetClassData['features'][$index] = [];
            }

            if (!isset($targetClassData['features'][$index][$feature])) {
                $targetClassData['features'][$index][$feature] = 0;
            }

            $targetClassData['features'][$index][$feature]++;
        }

        $targetClassData['count']++;
        $this->totalCount++;
    }

    public function classify($data) {
        $maxProb = PHP_INT_MIN;
        $bestClass = null;

        foreach ($this->classes as $targetClass => $targetClassData) {
            $classProb = log($targetClassData['count'] / $this->totalCount);

            foreach ($data as $index => $feature) {
                $featureData = $targetClassData['features'][$index] ?? [];
                $featureCount = $featureData[$feature] ?? 0;
                $totalFeatureCount = array_sum($featureData);

                $smoothedProb = ($featureCount + 1) / ($totalFeatureCount + count($featureData));

                $classProb += log($smoothedProb);
            }

            if ($classProb > $maxProb) {
                $maxProb = $classProb;
                $bestClass = $targetClass;
            }
        }

        return $bestClass;
    }
}
?>