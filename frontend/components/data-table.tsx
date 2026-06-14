type DataTableProps = {
  columns: string[];
  rows: Array<Array<string>>;
};

export function DataTable({ columns, rows }: DataTableProps) {
  return (
    <div className="overflow-hidden rounded-lg border border-[#d9e0ea] bg-white">
      <table className="w-full border-collapse text-left text-sm">
        <thead className="bg-[#f8fafc] text-[#4f5e75]">
          <tr>
            {columns.map((column) => (
              <th key={column} className="border-b border-[#d9e0ea] px-4 py-3 font-semibold">
                {column}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {rows.map((row, index) => (
            <tr key={index} className="border-b border-[#edf1f6] last:border-0">
              {row.map((cell, cellIndex) => (
                <td key={cellIndex} className="px-4 py-3 text-[#172033]">
                  {cell}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
