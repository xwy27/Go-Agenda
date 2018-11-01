# Go-Agenda

简单 CLI 会议管理系统

## 依赖

[Cobra](https://github.com/spf13/cobra#overview)

## 测试

Version1.0：静态测试通过

## 使用说明

1. Go Support

    安装 [Go 环境](https://golang.org/dl/)
1. Download Project

    使用如下命令进行项目下载：
    ```shell
    go get github.com/xwy27/Go-Agenda
    ```
    *如果下载失败，可手动下载：进入 $GOPATH/src/github.com 目录，利用 `git clone https://github.com/xwy27/Go-Agenda` 进行下载*
1. Compile & Run

    使用如下命令进行编译安装：
    ```shell
    go install github.com/xwy27/Go-Agenda
    ```
    然后，可以根据下文用法进行使用

## 功能

### register

- Usage:

      Go-Agenda register
        -u | --username Username
        -p | --password Password
        -e | --email Email
        -t | --telephone Phone_Number

- Description:

  注册新用户时，用户需设置一个唯一的用户名和一个密码。另外，还需登记邮箱及电话信息。

  如果注册时提供的用户名已由其他用户使用，反馈一个出错信息；成功注册后，反馈一个成功注册的信息。

### login

- Usage:

      Go-Agenda login
        -u | --username Username
        -p | --password Password

- Description:

  用户使用用户名和密码登录 Agenda 系统。

  用户名和密码同时正确则登录成功并反馈一个成功登录的信息。否则，登录失败并反馈一个失败登录的信息。

### logout

- Usage:

      Go-Agenda logout

- Description:

    已登录的用户登出系统后，只能使用用户注册和用户登录功能。

### listUsers

- Usage

      Go-Agenda listUsers

- Description:

    已登录的用户可以查看已注册的所有用户的用户名、邮箱及电话信息

### deleteUser

- Usage

      Go-Agenda deleteUser

- Description:

  已登录的用户可以删除本用户账户（即销号）。

  操作成功，需反馈一个成功注销的信息；否则，反馈一个失败注销的信息。

  以该用户为 发起者 的会议将被删除

  以该用户为 参与者 的会议将从 参与者 列表中移除该用户。若因此造成会议 参与者 人数为0，则会议也将被删除。

### createMeeting

- Usage:

      Go-Agenda createMeeting
        -t | --title Title
        -p | --participator Participator
        -s | --startTime yyyy-mm-ddThh:mm
        -e | --endTime yyyy-mm-ddThh:mm

- Description:

  已登录的用户可以添加一个新会议到其议程安排中。会议可以在多个已注册 用户间举行，不允许包含未注册用户。添加会议时提供的信息应包括：

  1. 会议主题(title)（在会议列表中具有唯一性）
  2. 会议参与者(participator)
  3. 会议起始时间(start time)
  4. 会议结束时间(end time)

  注意，任何用户都无法分身参加多个会议。如果用户已有的会议安排（作为发起者或参与者）与将要创建的会议在时间上重叠 （允许仅有端点重叠的情况），则无法创建该会议。

  用户应获得适当的反馈信息，以便得知是成功地创建了新会议，还是在创建过程中出现了某些错误。

### add / remove Participator

- Usage:

      Go-Agenda addParticipator | removeParticipator
        -t | --title Title
        -p | --Participator Participator

- Description:

  已登录的用户可以向 自己发起的某一会议增加/删除 参与者 。
  
  增加参与者时需要做 时间重叠 判断（允许仅有端点重叠的情况）。
  
  删除会议参与者后，若因此造成会议 参与者 人数为0，则会议也将被删除。

### queryMeetings

- Usage:

      Go-Agenda queryMeetings
        -s | --startTime yyyy-mm-ddThh:mm
        -e | --endTime yyyy-mm-ddThh:mm

- Description:

  已登录的用户可以查询自己的议程在某一时间段(time interval)内的所有会议安排。
  
  用户给出所关注时间段的起始时间和终止时间，返回该用户议程中在指定时间范围内找到的所有会议安排的列表。
  
  在列表中给出每一会议的起始时间、终止时间、主题、以及发起者和参与者。
  
  查询会议的结果应包括用户作为 发起者或参与者 的会议。

### cancelMeeting

- Usage:

      Go-Agenda cancelMeeting
        -t | --title Title

- Description

  已登录的用户可以取消 自己发起 的某一会议安排。
  
  取消会议时，需提供唯一标识：会议主题（title）。

### quitMeeting

- Usage:

      Go-Agenda quitMeeting
        -t | --title Title

- Description

  已登录的用户可以退出 自己参与 的某一会议安排。
  
  退出会议时，需提供一个唯一标识：会议主题（title）。若因此造成会议 参与者 人数为0，则会议也将被删除。

### clearMeeting

- Usage:

      Go-Agenda clearMeeting

- Description:
  
  已登录的用户可以清空 自己发起 的所有会议安排。