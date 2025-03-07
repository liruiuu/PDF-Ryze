import platform
import os
from pathlib import Path

if platform.system() == "Windows":
    logdir = Path(os.environ['USERPROFILE']) / ".pdf_ryze"
else:
    logdir = Path(os.environ['HOME']) / ".pdf_ryze"
logdir.mkdir(parents=True, exist_ok=True)
logpath = str(logdir / "pdf.log")
cmd_output_path = str(logdir / "cmd_output.json")
