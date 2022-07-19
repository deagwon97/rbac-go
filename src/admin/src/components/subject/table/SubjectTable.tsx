import { useState, useEffect } from 'react';
import axios from 'axios';
import Stack from '@mui/material/Stack';
import { styled } from '@mui/system';
import { Pagination } from '@mui/material';
import { Role } from 'components/role/row/RoleRow';
import SubjectRow from 'components/subject/row/SubjectRow';
import { API_URL } from 'App';

const grey = {
  50: '#F3F6F9',
  100: '#E7EBF0',
  200: '#E0E3E7',
  300: '#CDD2D7',
  400: '#B2BAC2',
  500: '#A0AAB4',
  600: '#6F7E8C',
  700: '#3E5060',
  800: '#2D3843',
  900: '#1A2027',
};

const Root = styled('div')(
  ({ theme }) => `
  table {
    font-family: IBM Plex Sans, sans-serif;
    font-size: 0.875rem;
    border-collapse: collapse;
    width: 100%;
  }

  td,
  th {
    border: 1px solid ${theme.palette.mode === 'dark' ? grey[800] : grey[200]};
    text-align: left;
    padding: 6px;
  }

  th {
    background-color: ${theme.palette.mode === 'dark' ? grey[900] : grey[100]};
  }
  `,
);

interface Subject {
  IsAllowed: boolean;
  SubjectID: number;
}

interface User {
  ID: number;
  Name: string;
}

interface SubjectTableProps {
  role: Role;
}

interface IDListForm {
  IDList: number[];
}

export default function SubjectTable({ role }: SubjectTableProps) {
  const [count, setCount] = useState<number>(0);
  const [subjectArray, setSubjectArray] = useState<Subject[]>([]);
  const [userArray, setUserArray] = useState<User[]>([]);
  const [pageNum, setPageNum] = useState<number>(1);
  const [pageCount, setPageCount] = useState<number>(1);
  const [readyRowData, SetReadyRowData] = useState<boolean>(false);
  const rowSize = 6;

  const getSubjectsOfRolePage = async (page: number) => {
    if (role !== null) {
      await axios
        .get(
          `${API_URL}/rbac/role/${role.ID}/subject?page=${page}&pageSize=${rowSize}`,
        )
        .then((res) => {
          setSubjectArray(res.data.Results);
          setCount(res.data.Count);
        });
    }
  };

  const getSubjectsName = async (subjectArray: Subject[]) => {
    const data: IDListForm = {
      IDList: [],
    };
    if (subjectArray.length > 0) {
      for (let idx = 0; idx < subjectArray.length; idx += 1) {
        data.IDList.push(subjectArray[idx].SubjectID);
      }
    }
    await axios.post(`${API_URL}/account/name/list`, data).then((res) => {
      setUserArray(res.data);
    });
  };

  useEffect(() => {
    getSubjectsOfRolePage(1);
  }, [role]);

  useEffect(() => {
    if (subjectArray !== []) {
      setPageCount(parseInt(`${(count + 1) / rowSize}`, 10));
    }
  }, [subjectArray]);

  useEffect(() => {
    if (userArray !== []) {
      getSubjectsName(subjectArray);
    }
  }, [subjectArray]);

  const handleChangePageNum = (
    event: React.ChangeEvent<unknown>,
    value: number,
  ) => {
    setPageNum(value);
    getSubjectsOfRolePage(value);
  };

  useEffect(() => {
    SetReadyRowData(
      subjectArray &&
        userArray &&
        subjectArray.length > 0 &&
        subjectArray.length === userArray.length,
    );
  }, [subjectArray, userArray]);

  return (
    <Root sx={{ width: 400, maxWidth: '100%' }}>
      {role && (
        <>
          <h1>Subjects Of Role</h1>
          <div style={{ minHeight: '435px' }}>
            <table aria-label="custom pagination table">
              <thead>
                <tr>
                  <th style={{ textAlign: 'center' }}>ID</th>
                  <th style={{ textAlign: 'center' }}>이름</th>
                  <th style={{ textAlign: 'center' }}>할당</th>
                </tr>
              </thead>
              <tbody>
                {readyRowData &&
                  subjectArray.map((subject: Subject, idx: number) => {
                    return (
                      <SubjectRow
                        key={subject.SubjectID}
                        subject={subject}
                        user={userArray[idx]}
                        roleID={role.ID}
                      />
                    );
                  })}
              </tbody>
            </table>
          </div>
          {readyRowData && (
            <Stack spacing={3}>
              <Pagination
                sx={{ margin: 'auto', marginTop: '10px' }}
                count={pageCount}
                defaultPage={pageNum}
                onChange={handleChangePageNum}
                shape="rounded"
              />
            </Stack>
          )}
        </>
      )}
    </Root>
  );
}
