import jsPDF from 'jspdf'
import html2canvas from 'html2canvas'

export async function exportarHorarioPDF(horario, bloques, escuelaNombre, periodoCodigo, serieDescripcion) {
  const grillaElement = document.getElementById('grilla-horario-export')
  if (!grillaElement) {
    throw new Error('No se encontró la grilla para exportar')
  }

  const canvas = await html2canvas(grillaElement, {
    scale: 2,
    useCORS: true,
    logging: false,
    backgroundColor: '#ffffff'
  })

  const imgData = canvas.toDataURL('image/png')
  const pdf = new jsPDF({
    orientation: 'landscape',
    unit: 'mm',
    format: 'a4'
  })

  const pageWidth = pdf.internal.pageSize.getWidth()
  const pageHeight = pdf.internal.pageSize.getHeight()
  const margin = 10
  const headerHeight = 25

  pdf.setFillColor(30, 41, 59)
  pdf.rect(0, 0, pageWidth, headerHeight, 'F')

  pdf.setTextColor(255, 255, 255)
  pdf.setFontSize(14)
  pdf.setFont('helvetica', 'bold')
  pdf.text(escuelaNombre, margin, 12)

  pdf.setFontSize(10)
  pdf.setFont('helvetica', 'normal')
  pdf.text(`Periodo: ${periodoCodigo}`, margin, 18)
  pdf.text(`Serie: ${serieDescripcion}`, margin + 70, 18)

  pdf.setFontSize(8)
  pdf.setFont('helvetica', 'italic')
  pdf.text(`Horario #${horario.id_horario} - Estado: ${horario.estado}`, margin, 23)

  const imgWidth = pageWidth - margin * 2
  const imgHeight = (canvas.height * imgWidth) / canvas.width

  let yPosition = headerHeight + 5

  if (imgHeight > pageHeight - headerHeight - margin - 10) {
    const ratio = (pageHeight - headerHeight - margin - 15) / imgHeight
    const scaledWidth = imgWidth * ratio
    const xOffset = (pageWidth - scaledWidth) / 2
    pdf.addImage(imgData, 'PNG', xOffset, yPosition, scaledWidth, pageHeight - headerHeight - margin - 15)
  } else {
    pdf.addImage(imgData, 'PNG', margin, yPosition, imgWidth, imgHeight)
  }

  pdf.setFontSize(9)
  pdf.setTextColor(80, 80, 80)

  pdf.setFillColor(59, 130, 246)
  pdf.rect(pageWidth - margin - 65, pageHeight - 7, 5, 5, 'F')
  pdf.setTextColor(59, 130, 246)
  pdf.text('TEORÍA', pageWidth - margin - 57, pageHeight - 3)

  pdf.setFillColor(34, 197, 94)
  pdf.rect(pageWidth - margin - 35, pageHeight - 7, 5, 5, 'F')
  pdf.setTextColor(34, 197, 94)
  pdf.text('PRÁCTICA', pageWidth - margin - 27, pageHeight - 3)

  const fileName = `${escuelaNombre.replace(/\s+/g, '_')}_${periodoCodigo.replace(/\s+/g, '_')}_${serieDescripcion.replace(/\s+/g, '_')}.pdf`
  pdf.save(fileName)
}
