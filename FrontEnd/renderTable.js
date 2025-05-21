document.getElementById('getAssignment').addEventListener('click', function () {
    dateInput = document.getElementById('calendar').value;

    if (!dateInput) {
        dateInput = getToday()
    }

    const url = `/api/getAssignment?date=${dateInput}`;

    fetch(url)
        .then(response => {
            if (!response.ok) {
                throw new Error('网络响应失败');
            }
            return response.json();
        })
        .then(data => {
            const responseDiv = document.getElementById('response');
            responseDiv.innerHTML = ''; // 清除旧内容

            const table = document.createElement('table');

            data.forEach(subArray => {
                const row = document.createElement('tr');

                subArray.forEach(item => {
                    const cell = document.createElement('td');
                    cell.textContent = item.Name || item.ID;

                    // 优先判断 Access 条件
                    if (item.Access <=3) {
                        cell.classList.add('cell_Moderator');
                    } else if (item.Note === 1) {
                        cell.classList.add('cell_SwitchOrRepay');
                    } else if (item.Note === 2) {
                        cell.classList.add('cell_Volunteering');
                    }

                    row.appendChild(cell);
                });

                table.appendChild(row);
            });
            const title =`<i><h5 align=center >${dateInput}网维值班表</h5></i>`
            const titleContainer = document.createElement('div');
            titleContainer.innerHTML = title
            responseDiv.appendChild(titleContainer)
            // 插入表格
            responseDiv.appendChild(table);

            // 添加图例说明
            const legendHTML = `
                <i class="table_notes"><span class="ZoneHead"></span>片区负责人<br></i>
                <i class="table_notes"><span class="Moderator"></span>管理层<br></i>
                <i class="table_notes"><span class="SwitchOrRepay"></span>换班/补班<br></i>
                <i class="table_notes"><span class="Volunteering"></span>蹭班<br></i>
            `;
            const legendContainer = document.createElement('div');
            legendContainer.innerHTML = legendHTML;
            responseDiv.appendChild(legendContainer);
        })
        .catch(error => {
            console.error('请求失败:', error);
            document.getElementById('response').innerHTML = '获取任务失败，请重试。';
        });
});

function getToday() {
  const today = new Date();

  const year = today.getFullYear();
  const month = String(today.getMonth() + 1).padStart(2, '0'); // 月份从0开始，需要+1
  const day = String(today.getDate()).padStart(2, '0');

  return `${year}-${month}-${day}`;
}







