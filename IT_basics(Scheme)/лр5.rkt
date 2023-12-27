(define feature-if-else #t)
(define feature-while-loop #t)
(define feature-repeat-loop #t)
(define feature-for-loop #t)
(define feature-break-continue #t)
(define feature-switch-case #t)
;(define feature-hi-level #t)
(define feature-tail-call #t)
(define feature-global #t)

(define (next x) (+ x 1))

(define (my-eval exprs)
  (eval exprs (interaction-environment)))


(define (find lexem program i)
  (if (= i (vector-length program))
      i
      (if (equal? (vector-ref program i) lexem)
          i
          (find lexem program (next i)))))
(define (watch x) (display x) x)

(define (findpare lexem1 lexem2 program j)
  (if (= j (vector-length program))
      #f
      (if (= (next (find lexem1 program j)) (find lexem2 program j))
          (find lexem2 program j)
          (findpare lexem1 lexem2 program (next j)))))

(define (set1 word x xs)
  (if (null? xs)
      '()
      (if (equal? (car word) (caar xs))
          (cons (list (car word) x) (set1 word x (cdr xs)))
          (cons (car xs) (set1 word x (cdr xs))))))


(define (minimum i j)
  (if (< i j)
      i
      j))
(define (minimum3 a b c)
  (minimum a (minimum b c)))

(define (interpret program start-stack)
  (let doit ((i 0) (stack start-stack) (stack-- '()) (slova '()) (constants '()))
    (if (= i (vector-length program))
        stack
        (let ((word (vector-ref program i)))
          (cond
            ((number? word) (doit (next i) (cons word stack) stack-- slova constants))
            ((equal? word '+) (doit (next i) (cons (+ (cadr stack) (car stack)) (cddr stack))
                                    stack-- slova constants))
            ((equal? word '-) (doit (next i) (cons (- (cadr stack) (car stack)) (cddr stack))
                                    stack-- slova constants))
            ((equal? word '*) (doit (next i) (cons (* (cadr stack) (car stack)) (cddr stack))
                                    stack-- slova constants))
            ((equal? word '/) (doit (next i) (cons (quotient (cadr stack) (car stack)) (cddr stack))
                                    stack-- slova constants))
            ((equal? word 'mod) (doit (next i) (cons (remainder (cadr stack) (car stack)) (cddr stack))
                                      stack-- slova constants))
            ((equal? word '>) (doit (next i) (cons (if (> (cadr stack) (car stack)) -1 0) (cddr stack))
                                    stack-- slova constants))
            ((equal? word '<) (doit (next i) (cons (if (< (cadr stack) (car stack)) -1 0) (cddr stack))
                                    stack-- slova constants))
            ((equal? word '=) (doit (next i) (cons (if (= (cadr stack) (car stack)) -1 0) (cddr stack))
                                    stack-- slova constants))
            ((equal? word 'and) (doit (next i) (cons (if (or (= (car stack) 0) (= (cadr stack) 0)) 0 -1)
                                                     (cddr stack))
                                      stack-- slova constants))
            ((equal? word 'or) (doit (next i) (cons (if (and (= (car stack) 0) (= (cadr stack) 0)) 0 -1)
                                                    (cddr stack))
                                     stack-- slova constants))
            ((equal? word 'neg) (doit (next i) (cons (- (car stack)) (cdr stack)) stack-- slova constants))
            ((equal? word 'not) (doit (next i) (cons (if (= (car stack) 0) -1 0) (cdr stack))
                                      stack-- slova constants))
            ((equal? word 'drop) (doit (next i) (cdr stack) stack-- slova constants))
            ((equal? word 'swap) (doit (next i) (append (list (cadr stack) (car stack)) (cddr stack))
                                       stack-- slova constants))
            ((equal? word 'dup) (doit (next i) (cons (car stack) stack) stack-- slova constants))
            ((equal? word 'over) (doit (next i) (cons (cadr stack) stack) stack-- slova constants))
            ((equal? word 'rot) (doit (next i)
                                      (append (list (caddr stack) (cadr stack) (car stack)) (cdddr stack))
                                      stack-- slova constants))
            ((equal? word 'depth) (doit (next i) (cons (length stack) stack) stack-- slova constants))
            ((equal? word 'define) (doit (next (find 'end program i)) stack stack--
                                         (cons (list (vector-ref program (next i)) (+ i 2)) slova)
                                         constants))
            ((or (equal? word 'end) (equal? word 'exit)) (doit (car stack--) stack (cdr stack--)
                                                               slova constants))
            ((equal? word 'if) (doit (if (zero? (car stack))
                                         (minimum (next (find 'else program i))
                                                  (next (find 'endif program i)))
                                         (next i))
                                     (cdr stack) stack-- slova constants))
            ((equal? word 'else) (doit (next (find 'endif program i)) stack stack-- slova constants))
            ((equal? word 'endif) (doit (next i) stack stack-- slova constants))
            ((equal? word 'while) (if (zero? (car stack))
                                      (doit (next (find 'wend program i)) (cdr stack)
                                            stack-- slova constants)
                                      (doit (next i) (cdr stack) (cons i stack--) slova constants)))
            ((equal? word 'wend) (doit (car stack--) stack (cdr stack--) slova constants))
            ((equal? word 'repeat) (doit (next i) stack (cons i stack--) slova constants))
            ((equal? word 'until) (doit (if (zero? (car stack)) (car stack--) (next i))
                                        (cdr stack) (cdr stack--) slova constants))
            ((equal? word 'for) (doit (next i) (cddr stack)
                                      (append (list (cadr stack) (car stack) (next i)) stack--)
                                      slova constants))
            ((equal? word 'i) (doit (next i) (cons (car stack--) stack) stack-- slova constants))
            ((equal? word 'next) (if (>= (car stack--) (cadr stack--))
                                     (doit (next i) stack (cdddr stack--) slova constants)
                                     (doit (caddr stack--) stack
                                           (cons (next (car stack--)) (cdr stack--))
                                           slova constants)))
            ((equal? word '&) (doit (+ i 2) (cons (cadr (assoc (vector-ref program (next i)) slova))
                                                  stack)
                                    stack-- slova constants))
            ((equal? word 'lam) (doit (next (find 'endlam program i)) (cons (next i) stack)
                                      stack-- slova constants))
            ((equal? word 'endlam) (doit (car stack--) stack (cdr stack--) slova constants))
            ((equal? word 'apply) (doit (car stack) (cdr stack) (cons (next i) stack--)
                                        slova constants))
            ((equal? word 'switch) (doit (if (findpare 'case (car stack) program 0)
                                             (- (findpare 'case (car stack) program 0) 1)
                                             (next (find 'endswitch program i)))
                                         (cdr stack) stack-- slova constants))
            ((equal? word 'case) (doit (+ i 2) stack stack-- slova constants))
            ((equal? word 'exitcase) (doit (next (find 'endswitch program i))
                                           stack stack-- slova constants))
            ((equal? word 'endswitch) (doit (next i) stack stack-- slova constants))
            ((equal? word 'break) (if (= (minimum3 (find 'wend program i)
                                                   (find 'next program i)
                                                   (find 'until program i))
                                         (find 'wend program i))
                                      (doit (next (find 'wend program i))
                                            stack stack-- slova constants)
                                      (if (= (minimum3 (find 'wend program i)
                                                       (find 'next program i)
                                                       (find 'until program i))
                                             (find 'next program i))
                                          (doit (next (find 'next program i))
                                                stack (cdddr stack--) slova constants)
                                          (doit (next (find 'until program i))
                                                stack (cdr stack--) slova constants))))
            ((equal? word 'continue) (if (= (minimum3 (find 'wend program i)
                                                      (find 'next program i)
                                                      (find 'until program i))
                                            (find 'wend program i))
                                         (doit (car stack--) stack (cdr stack--) slova constants)
                                         (if (= (minimum3 (find 'wend program i)
                                                          (find 'next program i)
                                                          (find 'until program i))
                                                (find 'next program i))
                                             (if (>= (car stack--) (cadr stack--))
                                                 (doit (next i) stack (cdddr stack--) slova constants)
                                                 (doit (caddr stack--) stack
                                                       (cons (next (car stack--)) (cdr stack--))
                                                       slova constants))
                                             (doit (find 'until program i)
                                                   stack stack-- slova constants))))
            ((equal? word 'tail) (doit (next (findpare 'define (vector-ref program (next i))
                                                       program 0))
                                       stack stack-- slova constants))
            ((equal? word 'defvar) (doit (+ i 3) stack stack-- slova
                                         (cons (list (vector-ref program (next i))
                                                     (vector-ref program (+ i 2))) constants)))
            ((equal? word 'set) (doit (+ i 2) (cdr stack) stack-- slova
                                      (set1 (assoc (vector-ref program (next i)) constants)
                                            (car stack) constants)))
            (else (if (assoc word slova)
                      (doit (cadr (assoc word slova)) stack (cons (next i) stack--)
                            slova constants)
                      (doit (next i) (cons (cadr (assoc word constants)) stack)
                            stack-- slova constants))))))))


;(interpret #(   define abs 
;                 dup 0 < 
;                 if neg endif 
;                 end 
;                 9 abs 
;                 -9 abs      ) (quote ()))
;
;(interpret #(   define =0? dup 0 = end
;                 define <0? dup 0 < end
;                 define signum
;                 =0? if exit endif
;                 <0? if drop -1 exit endif
;                 drop
;                 1
;                 end
;                 0 signum
;                 -5 signum
;                 10 signum       ) (quote ()))
;
;(interpret #(   define -- 1 - end
;                 define =0? dup 0 = end
;                 define =1? dup 1 = end
;                 define factorial
;                 =0? if drop 1 exit endif
;                 =1? if drop 1 exit endif
;                 dup --
;                 factorial
;                 *
;                 end
;                 0 factorial
;                 1 factorial
;                 2 factorial
;                 3 factorial
;                 4 factorial     ) (quote ()))
;
;(interpret #(   define =0? dup 0 = end
;                 define =1? dup 1 = end
;                 define -- 1 - end
;                 define fib
;                 =0? if drop 0 exit endif
;                 =1? if drop 1 exit endif
;                 -- dup
;                 -- fib
;                 swap fib
;                 +
;                 end
;                 define make-fib
;                 dup 0 < if drop exit endif
;                 dup fib
;                 swap --
;                 make-fib
;                 end
;                 10 make-fib     ) (quote ()))
;
;(interpret #(   define =0? dup 0 = end
;                 define gcd
;                 =0? if drop exit endif
;                 swap over mod
;                 gcd
;                 end
;                 90 99 gcd
;                 234 8100 gcd    ) '())