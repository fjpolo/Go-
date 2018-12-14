/***********************************************************************
 *
 *		Filename: maindefines.h
 *
 *
 *      Author: fpolo - fjpolo@gmail.com
 *
 *
 *  	Description:
 *
 *
 *  	Defined groups:
 *
 *						Priorities at which the tasks are created
 *						The number of items the queue can hold
 *						The periods assigned to the timers
 *						Mailboxes
 *
 *
 *
 * 		History:
 *
 *          12.07.2018:
 *          				- Creation
 *
 *          	.
 *          	.
 *          	.
 *          10.07.2018:
 *          				- Added comments
 *
 ***********************************************************************/

#ifndef MAINDEFINES_H_
#define MAINDEFINES_H_

/* Priorities at which the tasks are created. */
#define	configBLINKY_TASK_PRIORITY						( tskIDLE_PRIORITY + 3 )
#define	configLJULogo_TASK_PRIORITY						( tskIDLE_PRIORITY + 1 )
#define mainReceiveCANFrame_REQ_ID_TASK_PRIORITY		( tskIDLE_PRIORITY + 4 )
#define mainReceiveCANFrame_SEND_ID_TASK_PRIORITY		( tskIDLE_PRIORITY + 4 )
#define mainSendCANFrame_TASK_PRIORITY					( tskIDLE_PRIORITY + 3 )
#define mainuIP_TASK_PRIORITY							( tskIDLE_PRIORITY + 2 )

/* The rate at which data is sent to the queue, specified in milliseconds. */
#define mainQUEUE_SEND_FREQUENCY_MS			( 500 / portTICK_RATE_MS )

/* The number of items the queue can hold.  This is 1 as the receive task
will remove items as they are added so the send task should always find the
queue empty. */
#define mainQUEUE_LENGTH					( 1 )

/* Define an enumerated type used to identify the source of the data. */
typedef enum
{
	eSender1,
	eSender2
} DataSource_t;
/* Define the structure type that will be passed on the queue. */
typedef struct
{
	uint8_t ucValue;
	DataSource_t eDataSource;
} Data_t;
/* Declare two variables of type Data_t that will be passed on the queue. */
static const Data_t xStructsToSend[ 2 ] =
{
	{ 100, eSender1 }, /* Used by Sender1. */
	{ 200, eSender2 } /* Used by Sender2. */
};


/* Timers */
/* The periods assigned to the timers */
#define mainONE_SHOT_TIMER_PERIOD 			pdMS_TO_TICKS(100)
#define mainONE_SHOT_CLEAR_TIMER_PERIOD 	pdMS_TO_TICKS(1000)
#define mainTIMER_ONE_PERIOD				pdMS_TO_TICKS(10)
#define mainTIMER_TWO_PERIOD				pdMS_TO_TICKS(500)
#define mainTIMER_THREE_PERIOD				pdMS_TO_TICKS(5000)
#define mainSWITCH_READ_PERIOD				pdMS_TO_TICKS(100)
#define mainDSP_READ_PERIOD					pdMS_TO_TICKS(5)

/* Mailboxes used for demo. Keep mailboxes 4 apart if you want masks
independent - not affecting neighboring mailboxes. */
#define 	CANBOX_TX		    0x01	// Mailbox #1
#define 	CANBOX_RX 		    0x04	// Mailbox #4
#define 	CANBOX_RX_ASSIGNED	0x08	// Mailbox #8
#define 	CANBOX_RX_SENDID	0x14		// Mailbox #12
#define 	CANBOX_RX_REQID		0x10	// Mailbox #16

/******************************************************************************
Macro definitions
******************************************************************************/
/* TEST CAN ID */
#define 	TX_ID_DEFAULT 		0x001
#define 	RX_ID_DEFAULT 		0x001
#define 	REMOTE_TEST_ID		0x050


/* The WEB server uses string handling functions, which in turn use a bit more
stack than most of the other tasks. */
#define mainuIP_STACK_SIZE			( configMINIMAL_STACK_SIZE * 3 )


#endif /* MAINDEFINES_H_ */
