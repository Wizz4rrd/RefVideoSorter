import tkinter as tk
import tkinter.filedialog as fd
import tkinter.ttk as ttk
import os

window = tk.Tk()
window.title("Ref Video Converter")

button = tk.Button(window, text="EXIT APP", width=25, command=window.destroy)
button.pack()

window.geometry("400x600")
window.eval('tk::PlaceWindow . center')
window.attributes("-topmost",True)

dir_path = tk.StringVar()
lbl_path = tk.Label(window,textvariable=dir_path)
lbl_path.pack()

def get_directory():
    filepath = fd.askdirectory(initialdir="", title="Dialog box")
    dir_path.set(filepath)
    list_files(filepath)

def list_files(dir_path):
    for item in tree.get_children():
        tree.delete(item)
    for filename in os.listdir(dir_path):
        if os.path.isfile(os.path.join(dir_path, filename)):
            tree.insert("","end",values=filename,)
        


dialog_btn = tk.Button(window, text="Select directory", command=get_directory)
dialog_btn.pack()

tree = ttk.Treeview(window, columns=('Files'), show="headings", selectmode="browse")
tree.heading("Files",text="Files in directory")
tree.pack(padx=20,pady=20, fill="both", expand=True)

window.mainloop()