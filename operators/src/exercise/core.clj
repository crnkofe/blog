(ns exercise.core
  (:gen-class))

;; recursive one-liner because why not
(defn fibo-rec [n] (if (<= n 2) 1 (+ (fibo-rec (- n 1)) (fibo-rec (- n 2)))))

(def fibo-mem 
  "memoization stores results of previous calls to function with same arguments
   this is cached in-between consecutive calls to fibo-mem"
  (memoize
    (fn [n]
      (if (<= n 2) 1 (+ (fibo-mem (- n 2)) (fibo-mem (- n 1))))
    )
  )
)


(defn fibo-op 
  "Demonstrate ->> operator
  It works like a pipeline where each expression is fed into the next"
  [n]
  (if (<= n 2)
    1
    (loop [n-2 1 n-1 1 ct n]
      (if (<= ct 2)
        n-1
        (recur n-1 (->> [n-2 n-1] (apply +)) (dec ct))
      )
    )
  )
)

(defn -main
  "Operator fun."
  [& args]
  (loop [i 1]
    (when (< i 10)
      (println "Fibonnacci sequence(recursive:" (fibo-rec i))
      (println "Fibonnacci sequence(->>):" (fibo-op i))
      (println "Fibonnacci sequence(memoize):" (fibo-mem i))
      (recur (->> i inc))
    )
  )
)
