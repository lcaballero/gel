(defun tools-el/root-of (file-or-dir)
    (locate-dominating-file buffer-file-name file-or-dir))

(defun tools-el/dir ()
  (tools-el/root-of ".tools.el"))

(defun tools-el/file (name)
  (format "%s.tools.el/%s" (tools-el/dir) name))

(defun tools-el/load-all ()
  (interactive)
  (load (tools-el/file "build.el")))

(tools-el/load-all)


