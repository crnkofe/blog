(ns aoc2020-lessons.core
  (:gen-class))

(defn on-evaluate [x]
  (println x)
  (if (= x 5)
    (throw (Exception. "fail"))
    (+ x 5)
    )
  )

(defn demo-lazy-bug 
  "Throws exception later on after actually evaluating code"
  []
  (let [call-lazy-bug (map on-evaluate (range 10))]
    (println "no exception")
    (try 
      (doall call-lazy-bug)
      (catch Exception e
        (println "exception caught")
        )
      )
    )
  )

(defn -main
  "Demonstration of some lessons learned with doing AOC2020 in Clojure"
  [& args]
    (println "Lazy evaluation demo:")
    (demo-lazy-bug)
  )
