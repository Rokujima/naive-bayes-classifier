package main

import (
	"fmt"
	"math"
)

type NaiveBayesClassifier struct {
	classes    map[string]*classData
	totalCount int
}

type classData struct {
	count    int
	features map[int]map[string]int
}

func NewNaiveBayesClassifier() *NaiveBayesClassifier {
	return &NaiveBayesClassifier{
		classes:    make(map[string]*classData),
		totalCount: 0,
	}
}

func (nb *NaiveBayesClassifier) Train(data []string, targetClass string) {
	if _, ok := nb.classes[targetClass]; !ok {
		nb.classes[targetClass] = &classData{
			count:    0,
			features: make(map[int]map[string]int),
		}
	}

	targetClassData := nb.classes[targetClass]

	for featureIndex, feature := range data {
		if _, ok := targetClassData.features[featureIndex]; !ok {
			targetClassData.features[featureIndex] = make(map[string]int)
		}
		targetClassData.features[featureIndex][feature]++
	}

	targetClassData.count++
	nb.totalCount++
}

func (nb *NaiveBayesClassifier) Classify(data []string) string {
	maxProb := math.Inf(-1)
	var bestClass string

	for targetClass, targetClassData := range nb.classes {
		classProb := float64(targetClassData.count) / float64(nb.totalCount)

		for featureIndex, feature := range data {
			featureData := targetClassData.features[featureIndex]
			featureCount := featureData[feature]
			totalFeatureCount := 0
			for _, count := range featureData {
				totalFeatureCount += count
			}

			smoothedProb := float64(featureCount+1) / float64(totalFeatureCount+len(featureData))

			classProb *= smoothedProb
		}

		if classProb > maxProb {
			maxProb = classProb
			bestClass = targetClass
		}
	}

	return bestClass
}

func main() {
	//call classifier
	classifier := NewNaiveBayesClassifier()
}