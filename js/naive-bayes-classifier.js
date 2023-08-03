class NaiveBayesClassifier {
    constructor() {
      this.classes = {};
      this.totalCount = 0;
    }
  
    train(data, targetClass) {
      if (!this.classes[targetClass]) {
        this.classes[targetClass] = {
          count: 0,
          features: {},
        };
      }
  
      const targetClassData = this.classes[targetClass];
  
      data.forEach((feature, index) => {
        if (!targetClassData.features[index]) {
          targetClassData.features[index] = {};
        }
  
        if (!targetClassData.features[index][feature]) {
          targetClassData.features[index][feature] = 0;
        }
  
        targetClassData.features[index][feature]++;
      });
  
      targetClassData.count++;
      this.totalCount++;
    }
  
    classify(data) {
      let maxProb = -Infinity;
      let bestClass = null;
  
      for (const targetClass in this.classes) {
        const targetClassData = this.classes[targetClass];
        let classProb = Math.log(targetClassData.count / this.totalCount);
  
        data.forEach((feature, index) => {
          const featureData = targetClassData.features[index];
          const featureCount = featureData ? featureData[feature] || 0 : 0;
          const totalFeatureCount = Object.values(featureData || {}).reduce((acc, val) => acc + val, 0);
  
          const smoothedProb = (featureCount + 1) / (totalFeatureCount + Object.keys(featureData || {}).length);
  
          classProb += Math.log(smoothedProb);
        });
  
        if (classProb > maxProb) {
          maxProb = classProb;
          bestClass = targetClass;
        }
      }
  
      return bestClass;
    }
  }