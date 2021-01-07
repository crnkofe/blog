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
    (println "it would be nice to see a crash here")
    (try 
      (doall call-lazy-bug)
      (catch Exception e
        (println "but no")
        )
      )
    )
  )

(defn sequence-shuffler [n iterations]
  (loop [some-sequence (range n)
         steps 0]
      (if (>= steps iterations)
        some-sequence
        (recur (doall (concat (drop (/ n 4) some-sequence) (take (/ n 4) some-sequence))) (inc steps))
        )
      )
    )


(defn vector-shuffler [n iterations]
  (loop [some-vector (into [] (range n))
         steps 0]
      (if (>= steps iterations)
        some-vector
        (recur (into (subvec some-vector (/ n 4) n) (subvec some-vector 0 (/ n 4))) (inc steps))
        )
      )
    )

(defn sequence-vector-benchmark []
  (let [all-ns [1e1 1e2 1e3 1e4 1e5]
        steps 1e3]
    (loop [remaining-ns all-ns]
      (when (not (empty? remaining-ns))
        (println "n:" (first remaining-ns) "steps:" steps)
        (time (sequence-shuffler (first remaining-ns) steps))
        (time (vector-shuffler (first remaining-ns) steps))
        (recur (rest remaining-ns))
        )
      )
    )
  )

(defn -main
  "Demonstration of some lessons learned with doing AOC2020 in Clojure"
  [& args]
    (println "Lazy evaluation demo:")
    (demo-lazy-bug)
    (println "Sequence vs. Vectors:")
    (sequence-vector-benchmark)
  )
