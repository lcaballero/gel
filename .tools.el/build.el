(defun build/file (name args)
  (format "%s%s %s" (tools-el/root-of ".dir-locals.el") name args))

(defun build/sh (script func)
  (shell-command (build/file script func)))

(defun build/run-build ()
  (interactive)
  (build/sh "run.sh" "build"))

(defun build/run-clean ()
  (interactive)
  (build/sh "run.sh" "clean"))

(defun build/run-tests ()
  (interactive)
  (build/sh "run.sh" "tests"))

(defun build/run-cover ()
  (interactive)
  (build/sh "run.sh" "cover"))

(defun build/run-lint ()
  (interactive)
  (build/sh "run.sh" "lint"))

(defun build/run-gen ()
  (interactive)
  (build/sh "run.sh" "gen"))

(defhydra build (:color pink :hint nil :exit t)
  "
^Dev Local^
---------------------------
_t_: ./run.sh tests
_b_: ./run.sh build
_c_: ./run.sh clean

_v_: ./run.sh cover
_l_: ./run.sh lint
_g_: ./run.sh gen

"
  ;; local tooling for ./run.sh functions
  ("d" build/run-dist nil)
  ("b" build/run-build nil)
  ("c" build/run-clean nil)
  ("t" build/run-tests nil)
  ("l" build/run-lint nil)
  ("v" build/run-cover nil)
  ("g" build/run-gen nil)  

  ("q" nil "quit" :exit t))

(global-set-key (kbd "C-c v") 'build/body)
(global-set-key (kbd "C-c C-v") 'build/body)
