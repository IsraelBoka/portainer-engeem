import { columnHelper } from './helper';

export const className = columnHelper.accessor('ClassName', {
  header: 'Class Name',
  id: 'className',
  cell: ({ row }) => row.original.ClassName || '-',
});
