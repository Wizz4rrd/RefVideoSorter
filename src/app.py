import tkinter as tk
import tkinter.filedialog as fd

window = tk.Tk()
window.title("Ref Video Converter")

button = tk.Button(window, text="EXIT APP", width=25, command=window.destroy)
button.pack()

window.geometry("400x200")
window.eval('tk::PlaceWindow . center')
window.attributes("-topmost",True)

dir_path = tk.StringVar()
lbl_path = tk.Label(window,textvariable=dir_path)
lbl_path.pack()

def get_directory():
    filepath = fd.askdirectory(initialdir="", title="Dialog box")
    dir_path.set(filepath)



dialog_btn = tk.Button(window, text="Select directory", command=get_directory)
dialog_btn.pack()



window.mainloop()