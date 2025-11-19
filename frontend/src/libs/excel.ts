import * as XLSX from 'xlsx'

function exportTableToExcel(tableId: string, fileName: string = 'export.xlsx'): void {
  if (!fileName.endsWith('.xlsx')) {
    fileName += '.xlsx'
  }
  const table = document.getElementById(tableId)
  if (!table) throw new Error(`Table ${tableId} not found`)

  const workbook = XLSX.utils.table_to_book(table)
  XLSX.writeFile(workbook, fileName)
}

export default { exportTableToExcel }
