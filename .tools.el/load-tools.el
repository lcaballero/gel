(defun tools-el/dir ()
  (locate-dominating-file buffer-file-name ".tools.el"))

(defun tools-el/file (name)
  (format "%s.tools.el/%s" (tools-el/dir) name))

(defun tools-el/load-all ()
  (interactive)
  (load (tools-el/file "build.el")))

(tools-el/load-all)


