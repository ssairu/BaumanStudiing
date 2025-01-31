(define top 0)

(define-syntax use-assertations
  (syntax-rules ()
    ((use-assertations) (call-with-current-continuation (lambda (stack)
                                                          (set! top stack))))))

(define-syntax assert
  (syntax-rules ()
    ((assert expr) (if (not expr)
                       (begin (display "FAILED: ")
                              (display 'expr)
                              (newline)
                              (top))))))

(use-assertations)

(define (1/x x)
  (assert (not (zero? x)))
  (/ 1 x))


(define (count path)
  (let ((port (open-input-file path)))
    (if (eof-object? (peek-char port))
        0
        (if (or (equal? (peek-char port) #\return)(equal? (peek-char port) #\newline))
            (letrec ((func (lambda (x)
                             (let ((c (read-char x)))
                               (if (eof-object? c)
                                   0
                                   (if (equal? c #\newline)
                                       (if (not (or (equal? (peek-char x) #\return)
                                                    (equal? (peek-char x)#\newline)))
                                           (+ 1 (func x))
                                           (func x))
                                       (func x)))))))
              (- (func port) 1))
            (letrec ((func (lambda (x)
                             (let ((c (read-char x)))
                               (if (eof-object? c)
                                   0
                                   (if (equal? c #\newline)
                                       (if (not (or (equal? (peek-char x) #\return)
                                                    (equal? (peek-char x)#\newline)))
                                           (+ 1 (func x))
                                           (func x))
                                       (func x)))))))
              (func port))))))

(define (save-data data x)
  (let ((port (open-output-file x #:exists 'update)))
    (write data port)
    (close-output-port port)))

(define (load-data x)
  (let ((port (open-input-file x)))
    (let ((r (read port)))
      (close-input-port port)
      r)))


(define-syntax my-let
  (syntax-rules ()
    ((my-let ((name1 init1) ...) expr ...)
     ((lambda (name1 ...) expr ...) init1 ...))))

(define-syntax my-let*
  (syntax-rules ()
    ((my-let* () expr ...)
     (my-let () expr ...))
    ((my-let* ((name1 init1)) expr ...)
     (my-let ((name1 init1)) expr ...))
    ((my-let* ((name1 init1) (name2 init2) ...) expr ...)
     (my-let ((name1 init1))
             (my-let* ((name2 init2) ...) expr ...)))))


(define-syntax my-if
  (syntax-rules ()
    ((my-if cond? true-expr false-expr)
     (let ((true-prom (delay true-expr)) (false-prom (delay false-expr)))
       (force (or (and cond? true-prom) false-prom))))))


(define trib
  (let ((prev '((2 1) (1 0) (0 0))))
    (lambda (n)
      (let ((key (assoc n prev)))
        (if key
            (cadr key)
            (let ((curr (+ (trib (- n 1)) (trib (- n 2)) (trib (- n 3)))))
              (begin (set! prev (cons
                                 (list n curr)
                                 prev))
                     curr)))))))

(define-syntax when
  (syntax-rules ()
    ((when cond? expr1 ...)
     (if cond? (begin expr1 ...)))))

(define-syntax unless
  (syntax-rules ()
    ((unless cond? expr1 ...)
     (if (not cond?) (begin expr1 ...)))))

(define-syntax for
  (syntax-rules (in as)
    ((for i in xs expr1 ...)
     (letrec ((count (lambda (vals)
                       (if (not (null? vals))
                           (let ((i (car vals)))
                             expr1 ...
                             (count (cdr vals)))))))
       (count xs)))
    ((for xs as i expr1 ...)
     (for i in xs expr1 ...))))


(define-syntax while
  (syntax-rules ()
    ((while cond? expr1 ...)
     (letrec ((count (lambda ()
                       (if cond?
                           (begin expr1 ... (count))))))
       (count)))))

(define-syntax repeat
  (syntax-rules (until)
    ((repeat (expr1 ...) until cond?)
     (letrec ((count (lambda ()
                       expr1 ...
                       (if (not cond?)
                           (count)))))
       (count)))))


(define-syntax cout
  (syntax-rules (<< endl)
    ((cout << endl)
     (newline))
    ((cout << endl . expressions)
     (begin (newline)
            (cout . expressions)))
    ((cout << expr1)
     (display expr1))
    ((cout << expr1 . expressions)
     (begin (display expr1)
            (cout . expressions)))))
