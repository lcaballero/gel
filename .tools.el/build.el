(defun build/root ()
  (locate-dominating-file buffer-file-name ".dir-locals.el"))

(defun build/file (name args)
  (format "%s%s %s" (build/root) name args))

(defun build/sh (script func)
  (shell-command (build/file script func)))

(defun build/run-dist ()
  (interactive)
  (build/sh "run.sh" "dist"))

(defun build/run-build ()
  (interactive)
  (build/sh "run.sh" "build"))

(defun build/run-clean ()
  (interactive)
  (build/sh "run.sh" "clean"))

(defun build/run-tests ()
  (interactive)
  (build/sh "run.sh" "tests"))

(defun build/run-lint ()
  (interactive)
  (build/run "run.sh" "lint"))

(defhydra build (:color pink :hint nil :exit t)
  "
^Dev Local^
---------------------------
_d_: ./run.sh dist
_b_: ./run.sh build
_c_: ./run.sh clean
_t_: ./run.sh tests
_l_: ./run.sh lint
"
  ;; local tooling for ./run.sh functions
  ("d" build/run-dist nil)
  ("b" build/run-build nil)
  ("c" build/run-clean nil)
  ("t" build/run-tests nil)
  ("l" build/run-lint nil)

  ("q" nil "quit" :exit t))

(global-set-key (kbd "C-c v") 'build/body)
(global-set-key (kbd "C-c C-v") 'build/body)
