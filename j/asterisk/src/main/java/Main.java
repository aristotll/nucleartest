import org.asteriskjava.live.AsteriskServerListener;
import org.asteriskjava.manager.*;
import org.asteriskjava.manager.action.OriginateAction;
import org.asteriskjava.manager.event.ChannelReloadEvent;
import org.asteriskjava.manager.event.ManagerEvent;
import org.asteriskjava.manager.event.NewStateEvent;
import org.asteriskjava.manager.event.StatusEvent;
import org.asteriskjava.manager.response.ManagerResponse;

import java.io.IOException;
import java.util.Map;

public class Main implements AsteriskServerListener {
    private ManagerConnection managerConnection;

    public Main() throws IOException {
        ManagerConnectionFactory factory = new ManagerConnectionFactory(
                "123.113.44.218",
                50380,
                "admin",
                "freepbx_amp111");

        this.managerConnection = factory.createManagerConnection();
    }

    public void run() throws IOException, AuthenticationFailedException,
            TimeoutException {
        OriginateAction originateAction;
        ManagerResponse originateResponse;

        originateAction = new OriginateAction();
        originateAction.setChannel("SIP/801");
        originateAction.setContext("from-internal");
        originateAction.setExten("804");
        originateAction.setPriority(1);
        originateAction.setTimeout(new Integer(30000));
        originateAction.setCallerId("801");
        originateAction.setAsync(true);

        // connect to Asterisk and log in
        managerConnection.login();

        managerConnection.addEventListener((event) -> {
            System.out.println(event);
        });
        // send the originate action and wait for a maximum of 30 seconds for Asterisk
        // to send a reply
        originateResponse = managerConnection.sendAction(originateAction, 30000);
        //Map<String, Object> map = originateResponse.getAttributes();
        //System.out.println(map);

        // print out whether the originate succeeded or not
        //System.out.println(originateResponse.getEventList());

        // and finally log off and disconnect
        managerConnection.logoff();
    }

    public static void main(String[] args) throws Exception {
        Main m = new Main();
        m.run();
    }

//    @Override
//    public void onManagerEvent(ManagerEvent managerEvent) {
//        System.out.println(managerEvent);
//    }
}
