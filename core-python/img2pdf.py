# Required packages:
# reportlab

from reportlab.pdfgen import canvas
from reportlab.lib.utils import ImageReader
import os, os.path, sys

def createPDF(pdfname, files):
	'''Create PDF file from a list.
	pdfname: a string path for the PDF file.
	files: a list with each element containing a string path for an image file.'''
	pdf = canvas.Canvas(pdfname)
	i = 1
	for file in files:
		img = ImageReader(file)
		pdf.setPageSize(img.getSize())
		try:
			pdf.drawImage(img, 0, 0)
		except OSError:
			pdf.drawImage(file, 0, 0)
		pdf.showPage()
		print('{}: {}'.format(i, os.path.basename(file)))
		i += 1
	pdf.save()

def getImages(path):
	'''Get images in a directory and generate absolute paths.'''
	files = os.listdir(path)
	files.sort()
	return [os.path.join(path, i) for i in files if i.lower().endswith('.jpg') or i.lower().endswith('.png')]

def getPDFName(path):
	'''Generate PDF file name from the directory. PDF will have the same name as the directory and ".pdf" extension.'''
	return os.path.join(os.path.dirname(path), os.path.basename(os.path.normpath(path)) + '.pdf')

def main():
	if len(sys.argv) > 1:
		for directory in sys.argv[1:]:
			images = getImages(directory)
			if len(images) > 0:
				pdf = getPDFName(directory)
				createPDF(pdf, images)
				print('PDF file created: {} with {} image file(s).'.format(os.path.basename(pdf), len(images)))
		input('Press any key to continue...') # Pause
	else:
		print('Usage: ' + sys.argv[0] + ' <directory that contains images> [<directory2 that contains images> ...]\nYou can drag and drop folders onto the script.')

if __name__ == '__main__':
	main()
