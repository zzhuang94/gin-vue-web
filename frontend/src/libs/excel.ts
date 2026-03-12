import * as XLSX from 'xlsx'

const DEFAULT_ID = 'index-table-id'

function exportExcel(fileName: string = 'export.xlsx', tableId: string = DEFAULT_ID) {
  if (!fileName.endsWith('.xlsx')) {
    fileName += '.xlsx'
  }
  const table = document.getElementById(tableId)
  if (!table) throw new Error(`Table ${tableId} not found`)

  const workbook = XLSX.utils.table_to_book(table)
  XLSX.writeFile(workbook, fileName)
}

export default { DEFAULT_ID, exportExcel }
