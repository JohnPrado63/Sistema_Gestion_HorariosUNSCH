import jsPDF from 'jspdf'

function getSerieTexto(serieNumero, semestre) {
  const cicloBase = Math.floor(serieNumero / 100)
  const cicloNum = (cicloBase - 1) * 2 + (semestre === 'II' ? 2 : 1)
  const ciclosRomanos = ['I', 'II', 'III', 'IV', 'V', 'VI', 'VII', 'VIII', 'IX', 'X']
  const cicloRomano = ciclosRomanos[cicloNum - 1]
  return `${serieNumero}-${semestre} (Ciclo ${cicloRomano})`
}

function addHeader(pdf, escuelaNombre, periodoCodigo, serieTexto, horarioId, estado) {
  const pageWidth = pdf.internal.pageSize.getWidth()
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
  pdf.text(`Serie: ${serieTexto}`, margin + 70, 18)

  pdf.setFontSize(8)
  pdf.setFont('helvetica', 'italic')
  pdf.text(`Horario #${horarioId} - Estado: ${estado}`, margin, 23)
}

function addLegend(pdf) {
  const pageWidth = pdf.internal.pageSize.getWidth()
  const pageHeight = pdf.internal.pageSize.getHeight()
  const margin = 10

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
}

function addGrillaImage(pdf, imgData) {
  const pageWidth = pdf.internal.pageSize.getWidth()
  const pageHeight = pdf.internal.pageSize.getHeight()
  const margin = 10
  const headerHeight = 25

  const imgWidth = pageWidth - margin * 2
  const imgHeight = (imgHeight => {
    const canvas = document.createElement('canvas')
    return imgHeight
  })(0)

  let yPosition = headerHeight + 5

  if (imgHeight > pageHeight - headerHeight - margin - 10 || !imgHeight) {
    const ratio = (pageHeight - headerHeight - margin - 15) / 200
    const scaledWidth = imgWidth * ratio
    const xOffset = (pageWidth - scaledWidth) / 2
    pdf.addImage(imgData, 'PNG', xOffset, yPosition, scaledWidth, pageHeight - headerHeight - margin - 15)
  } else {
    pdf.addImage(imgData, 'PNG', margin, yPosition, imgWidth, imgHeight)
  }
}

export function generarPDFConImagenes(paginas, escuelaNombre, periodoCodigo) {
  const pdf = new jsPDF({
    orientation: 'landscape',
    unit: 'mm',
    format: 'a4'
  })

  const pageWidth = pdf.internal.pageSize.getWidth()
  const pageHeight = pdf.internal.pageSize.getHeight()
  const margin = 10
  const headerHeight = 25

  paginas.forEach((pagina, i) => {
    if (i > 0) {
      pdf.addPage()
    }

    const serieTexto = getSerieTexto(pagina.serieNumero, pagina.semestre)

    pdf.setFillColor(30, 41, 59)
    pdf.rect(0, 0, pageWidth, headerHeight, 'F')

    pdf.setTextColor(255, 255, 255)
    pdf.setFontSize(14)
    pdf.setFont('helvetica', 'bold')
    pdf.text(escuelaNombre, margin, 12)

    pdf.setFontSize(10)
    pdf.setFont('helvetica', 'normal')
    pdf.text(`Periodo: ${periodoCodigo}`, margin, 18)
    pdf.text(`Serie: ${serieTexto}`, margin + 70, 18)

    pdf.setFontSize(8)
    pdf.setFont('helvetica', 'italic')
    pdf.text(`Horario #${pagina.horarioId} - Estado: ${pagina.estado}`, margin, 23)

    if (pagina.imgData) {
      const imgWidth = pageWidth - margin * 2
      const yPosition = headerHeight + 5

      pdf.addImage(pagina.imgData, 'PNG', margin, yPosition, imgWidth, 140)
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
  })

  const fileName = `${escuelaNombre.replace(/\s+/g, '_')}_${periodoCodigo.replace(/\s+/g, '_')}_TODOS.pdf`
  pdf.save(fileName)
}
